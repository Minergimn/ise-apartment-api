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
	Lock(ctx context.Context, n uint64) ([]apartment.ApartmentEvent, error)
	Unlock(ctx context.Context, eventIDs []uint64) error

	Add(ctx context.Context, events []apartment.ApartmentEvent) error
	Remove(ctx context.Context, eventIDs []uint64) error
}

type eventRepo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewEventRepo returns EventRepo interface
func NewEventRepo(db *sqlx.DB, batchSize uint) EventRepo {
	return &eventRepo{db: db, batchSize: batchSize}
}

func pgQb() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

func (e eventRepo) Lock(ctx context.Context, n uint64) ([]apartment.ApartmentEvent, error) {
	query := pgQb().Select("*").From("apartments_events").
		Where(sq.Eq{"is_locked": false}).
		Where(sq.Eq{"is_deleted": false}).
		Where(sq.Eq{"status": apartment.Deferred.String()}).
		OrderBy("id").Limit(n)

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res []apartment.ApartmentEvent
	err = e.db.SelectContext(ctx, &res, s, args...)
	if err != nil {
		return nil, err
	}

	var eventIDs []uint64
	for _, apt := range res {
		eventIDs = append(eventIDs, apt.ID)
	}

	err = e.setLocked(nil, eventIDs, true)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (e eventRepo) Unlock(ctx context.Context, eventIDs []uint64) error {
	return e.setLocked(ctx, eventIDs, false)
}

func (e eventRepo) Add(ctx context.Context, events []apartment.ApartmentEvent) error {
	query := pgQb().Insert("apartments_events").Columns(
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

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}

	var id uint64
	err = e.db.GetContext(ctx, &id, s, args...)
	if err != nil {
		return err
	}

	if id != 0 {
		return nil
	} else {
		return sql.ErrNoRows
	}
}

func (e eventRepo) Remove(ctx context.Context, eventIDs []uint64) error {
	query := pgQb().Update("apartments_events").
		Set("is_deleted", 1).
		Where(sq.Eq{"id": eventIDs})

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = e.db.ExecContext(ctx, s, args...)
	return err
}

func (e eventRepo) setLocked(ctx context.Context, eventIDs []uint64, value bool) error {
	query := pgQb().Update("apartments_events").
		Set("is_locked", value).
		Where(sq.Eq{"id": eventIDs})

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = e.db.ExecContext(ctx, s, args...)
	return err
}
