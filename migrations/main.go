package main

import (
	"embed"

	"github.com/ozonmp/ise-apartment-api/internal/config"
	db "github.com/ozonmp/ise-apartment-api/internal/database"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	conn, err := db.NewPostgres(cfg.Database.DSN, cfg.Database.Driver)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open() error")
	}
	defer conn.Close()

	goose.SetBaseFS(embedMigrations)

	const cmd = "up"

	err = goose.Run(cmd, conn.DB, "migrations")
	if err != nil {
		log.Fatal().Err(err).Msg("goose.Status() error")
	}

}