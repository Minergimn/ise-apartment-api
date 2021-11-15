package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ozonmp/ise-apartment-api/internal/logger"
	"github.com/ozonmp/ise-apartment-api/internal/tracer"
	"github.com/pressly/goose/v3"
	gelf "github.com/snovichkov/zap-gelf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/ozonmp/ise-apartment-api/internal/config"
	"github.com/ozonmp/ise-apartment-api/internal/database"
	"github.com/ozonmp/ise-apartment-api/internal/server"
)

var (
	batchSize uint = 2
)

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.ErrorKV(ctx, "Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	syncLogger := initLogger(ctx, cfg)
	defer syncLogger()

	tracing, err := tracer.NewTracer(ctx, &cfg)
	if err != nil {
		logger.ErrorKV(ctx, "Failed init tracing")

		return
	}
	defer tracing.Close()

	migration := flag.Bool("migration", true, "Defines the migration start option")
	flag.Parse()

	logger.InfoKV(ctx, fmt.Sprintf("Starting service: %s", cfg.Project.Name),
		"version", cfg.Project.Version,
		"commitHash", cfg.Project.CommitHash,
		"debug", cfg.Project.Debug,
		"environment", cfg.Project.Environment,
	)

	dsn := cfg.Database.DSN

	db, err := database.NewPostgres(ctx, dsn, cfg.Database.Driver)
	if err != nil {
		logger.FatalKV(ctx, "Failed init postgres")
	}
	defer db.Close()

	*migration = true
	if *migration {
		if err = goose.Up(db.DB, cfg.Database.Migrations); err != nil {
			logger.ErrorKV(ctx, "Migration failed")

			return
		}
	}

	if err := server.NewGrpcServer(db, batchSize).Start(&cfg); err != nil {
		logger.ErrorKV(ctx, "Failed creating gRPC server")

		return
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
		notSugaredLogger.Sync()
	}
}
