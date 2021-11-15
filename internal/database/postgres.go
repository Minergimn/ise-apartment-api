package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/ise-apartment-api/internal/logger"
)

// NewPostgres returns DB
func NewPostgres(ctx context.Context, dsn, driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		logger.ErrorKV(ctx, "failed to create database connection")

		return nil, err
	}

	// need to uncomment for homework-4
	if err = db.Ping(); err != nil {
		logger.ErrorKV(ctx, "failed ping the database")

		return nil, err
	}

	return db, nil
}
