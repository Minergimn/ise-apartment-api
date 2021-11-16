package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/opentracing/opentracing-go"

	model "github.com/ozonmp/ise-apartment-api/internal/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// Repo is DAO for Apartment
type Repo interface {
	GetApartment(ctx context.Context, apartmentID uint64) (*model.Apartment, error)
	ListApartments(ctx context.Context, cursor uint64, limit uint64, owner string, object string, ids []uint64) ([]model.Apartment, error)
	CreateApartment(ctx context.Context, apartment *model.Apartment) (uint64, error)
	DeleteApartment(ctx context.Context, apartmentID uint64) (bool, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func pgQb() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

func (r *repo) GetApartment(ctx context.Context, apartmentID uint64) (*model.Apartment, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.GetApartment")
	defer span.Finish()

	query := pgQb().Select("*").From("apartments").Where(sq.Eq{"id": apartmentID})

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res model.Apartment
	err = r.db.GetContext(ctx, &res, s, args...)

	if res.Status == model.Deleted {
		return nil, errors.New(fmt.Sprintf("Apartment id %d found but in deleted status", apartmentID))
	}

	return &res, err
}

func (r *repo) ListApartments(ctx context.Context, cursor uint64, limit uint64, owner string, object string, ids []uint64) ([]model.Apartment, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.ListApartments")
	defer span.Finish()

	query := pgQb().Select("*").From("apartments").
		Where(sq.Eq{"status": 0}).
		OrderBy("id")

	if cursor > 0 {
		query = query.Offset(cursor)
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	if owner != "" {
		query = query.Where(sq.Eq{"owner": owner})
	}

	if object != "" {
		query = query.Where(sq.Eq{"object": object})
	}

	if ids != nil {
		query = query.Where(sq.Eq{"id": ids})
	}

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res []model.Apartment
	err = r.db.SelectContext(ctx, &res, s, args...)

	return res, err
}

func (r *repo) CreateApartment(ctx context.Context, apartment *model.Apartment) (uint64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.CreateApartment")
	defer span.Finish()

	query := pgQb().Insert("apartments").
		Columns("object", "owner", "status").
		Values(apartment.Object, apartment.Owner, 0).
		Suffix("RETURNING id").
		RunWith(r.db)

	s, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}
	var id uint64
	err = r.db.GetContext(ctx, &id, s, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) DeleteApartment(ctx context.Context, apartmentID uint64) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.DeleteApartment")
	defer span.Finish()

	apartment, err := r.GetApartment(ctx, apartmentID)
	if err != nil {
		return false, err
	}

	if apartment == nil {
		return false, nil
	}

	if apartment.Status == model.Deleted {
		return true, nil
	}

	query := pgQb().Update("apartments").
		Set("status", 1).
		Where(sq.Eq{"id": apartmentID})

	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}
	_, err = r.db.ExecContext(ctx, s, args...)
	return true, err
}
