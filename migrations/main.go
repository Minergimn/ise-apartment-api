package main

import (
	"context"
	"embed"
	"github.com/ozonmp/ise-apartment-api/internal/logger"
	gelf "github.com/snovichkov/zap-gelf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"

	"github.com/ozonmp/ise-apartment-api/internal/config"
	db "github.com/ozonmp/ise-apartment-api/internal/database"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	ctx := context.Background()
	cfg := config.GetConfigInstance()

	syncLogger := initLogger(ctx, cfg)
	defer syncLogger()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.ErrorKV(ctx, "Failed init configuration")
	}

	conn, err := db.NewPostgres(ctx, cfg.Database.DSN, cfg.Database.Driver)
	if err != nil {
		logger.ErrorKV(ctx, "sql.Open() error")
	}
	defer conn.Close()

	goose.SetBaseFS(embedMigrations)

	const cmd = "up"

	err = goose.Run(cmd, conn.DB, "migrations")
	if err != nil {
		logger.ErrorKV(ctx, "goose.Status() error")
	}

}

func initLogger(ctx context.Context, cfg config.Config) (syncFn func()) {
	loggingLevel := zap.InfoLevel
	if cfg.Project.Debug {
		loggingLevel = zap.DebugLevel
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zap.NewAtomicLevelAt(loggingLevel),
	)

	gelfCore, err := gelf.NewCore(
		gelf.Addr(cfg.Telemetry.GraylogPath),
		gelf.Level(loggingLevel),
	)
	if err != nil {
		logger.FatalKV(ctx, "sql.Open() error", "err", err)
	}

	notSugaredLogger := zap.New(zapcore.NewTee(consoleCore, gelfCore))

	sugaredLogger := notSugaredLogger.Sugar()
	logger.SetLogger(sugaredLogger.With(
		"service", cfg.Project.ServiceName,
	))

	return func() {
		notSugaredLogger.Sync() //nolint:errcheck
	}
}
