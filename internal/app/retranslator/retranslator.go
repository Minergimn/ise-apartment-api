package retranslator

import (
	"context"
	apartment "github.com/ozonmp/ise-apartment-api/internal/model"
	"time"

	"github.com/ozonmp/ise-apartment-api/internal/app/consumer"
	"github.com/ozonmp/ise-apartment-api/internal/app/producer"
	"github.com/ozonmp/ise-apartment-api/internal/app/repo"
	"github.com/ozonmp/ise-apartment-api/internal/app/sender"

	"github.com/gammazero/workerpool"
)

//Retranslator comment for linter
type Retranslator interface {
	Start(ctx context.Context)
	Close(ctx context.Context)
}

//Config comment for linter
type Config struct {
	ChannelSize uint64

	ConsumerCount  uint64
	ConsumeSize    uint64
	ConsumeTimeout time.Duration

	ProducerCount uint64
	WorkerCount   int

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events     chan apartment.ApartmentEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
}

//NewRetranslator comment for linter
func NewRetranslator(cfg Config) Retranslator {
	events := make(chan apartment.ApartmentEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumer := consumer.NewDbConsumer(
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		cfg.Repo,
		events)
	producer := producer.NewKafkaProducer(
		cfg.ProducerCount,
		cfg.Sender,
		events,
		cfg.Repo,
		workerPool)

	return &retranslator{
		events:     events,
		consumer:   consumer,
		producer:   producer,
		workerPool: workerPool,
	}
}

func (r *retranslator) Start(ctx context.Context) {
	r.producer.Start(ctx)
	r.consumer.Start(ctx)
}

func (r *retranslator) Close(ctx context.Context) {
	r.consumer.Close(ctx)
	r.producer.Close(ctx)
	r.workerPool.StopWait()
}
