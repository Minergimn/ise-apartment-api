package repo

import (
	"context"
	"database/sql"
	"time"

	apartment "github.com/ozonmp/ise-apartment-api/internal/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type EventRepo interface {
	Lock(n uint64) ([]apartment.ApartmentEvent, error)
	Unlock(eventIDs []uint64) error

	Add(events []apartment.ApartmentEvent) error
	Remove(eventIDs []uint64) error
}

type eventRepo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewEventRepo returns EventRepo interface
func NewEventRepo(db *sqlx.DB, batchSize uint) EventRepo {
	return &eventRepo{db: db, batchSize: batchSize}
}

func (e eventRepo) Lock(n uint64) ([]apartment.ApartmentEvent, error) {
	query := sq.Select("*").PlaceholderFormat(sq.Dollar).From("apartments_events")
	query = query.Where(sq.Eq{"is_locked": false})
	query = query.Where(sq.Eq{"is_deleted": false})
	query = query.Where(sq.Eq{"status": apartment.Deferred.String()})
	query = query.OrderBy("id").Limit(n)

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res []apartment.ApartmentEvent
	err = e.db.SelectContext(context.Background(), &res, s, args...)
	if err != nil {
		return nil, err
	}

	var eventIDs []uint64
	for _, apt := range res {
		eventIDs = append(eventIDs, apt.ID)
	}

	err = e.setLocked(eventIDs, true)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (e eventRepo) Unlock(eventIDs []uint64) error {
	return e.setLocked(eventIDs, false)
}

func (e eventRepo) Add(events []apartment.ApartmentEvent) error {
	query := sq.Insert("apartments_events").PlaceholderFormat(sq.Dollar).Columns(
		"apartment_id",
		"type",
		"status",
		"payload",
		"is_deleted",
		"is_locked",
		"updated")

	for _, event := range events {
		query = query.Values(event.ApartmentId, event.Type.String(), event.Status.String(), event.Entity, false, false, time.Now())
	}

	query = query.Suffix("RETURNING id").RunWith(e.db)

	rows, err := query.QueryContext(context.Background())
	if err != nil {
		return err
	}

	var id uint64
	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return err
		}

		return nil
	} else {
		return sql.ErrNoRows
	}
}

func (e eventRepo) Remove(eventIDs []uint64) error {
	query := sq.Update("apartments_events").PlaceholderFormat(sq.Dollar).Set("is_deleted", 1).Where(sq.Eq{"id": eventIDs})
	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = e.db.ExecContext(context.Background(), s, args...)
	return err
}

func (e eventRepo) setLocked(eventIDs []uint64, value bool) error {
	query := sq.Update("apartments_events").PlaceholderFormat(sq.Dollar).Set("is_locked", value).Where(sq.Eq{"id": eventIDs})
	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = e.db.ExecContext(context.Background(), s, args...)
	return err
}
