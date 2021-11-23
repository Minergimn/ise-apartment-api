package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ozonmp/ise-apartment-api/internal/app/retranslator"
)

func main() {

	sigs := make(chan os.Signal, 1)

	cxt := context.Background()

	cfg := retranslator.Config{
		ChannelSize:   512,
		ConsumerCount: 2,
		ConsumeSize:   10,
		ProducerCount: 28,
		WorkerCount:   2,
		Ctx:           cxt,
	}

	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
