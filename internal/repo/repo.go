package repo

import (
	"context"

	model "github.com/ozonmp/ise-apartment-api/internal/model"

	"github.com/jmoiron/sqlx"
)

// Repo is DAO for Apartment
type Repo interface {
	DescribeApartment(ctx context.Context, apartmentID uint64) (*model.Apartment, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeApartment(ctx context.Context, apartmentID uint64) (*model.Apartment, error) {
	return nil, nil
}
