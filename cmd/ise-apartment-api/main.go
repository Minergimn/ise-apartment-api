package main

import (
	"context"
	event "github.com/ozonmp/ise-apartment-api/internal/app/repo"
	"github.com/ozonmp/ise-apartment-api/internal/app/sender"
	"github.com/ozonmp/ise-apartment-api/internal/config"
	"github.com/ozonmp/ise-apartment-api/internal/database"
	"github.com/ozonmp/ise-apartment-api/internal/logger"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/ozonmp/ise-apartment-api/internal/app/retranslator"
)

func main() {

	sigs := make(chan os.Signal, 1)

	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.ErrorKV(ctx, "Failed init configuration")
	}
	mainGfg := config.GetConfigInstance()

	sender, err := sender.NewEventSender(mainGfg.Kafka.Brokers)
	if err != nil {
		logger.FatalKV(ctx, "failed to create EventSender")
	}

	db, err := database.NewPostgres(ctx, mainGfg.Database.DSN, mainGfg.Database.Driver)
	if err != nil {
		logger.FatalKV(ctx, "Failed init postgres")
	}
	defer db.Close()
	repo := event.NewEventRepo(db, 2)

	cfg := retranslator.Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: 10,
		ProducerCount:  28,
		WorkerCount:    2,
		Ctx:            ctx,
		Sender:         sender,
		SenderTopic:    mainGfg.Kafka.Topic,
		Repo:           repo,
	}

	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
