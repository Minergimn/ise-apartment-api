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
	Start()
	Close()
}

//Config comment for linter
type Config struct {
	ChannelSize uint64

	ConsumerCount  uint64
	ConsumeSize    uint64
	ConsumeTimeout time.Duration

	ProducerCount uint64
	WorkerCount   int

	Ctx context.Context

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events     chan apartment.Event
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
	ctx        context.Context
	cancelFunc context.CancelFunc
}

//NewRetranslator comment for linter
func NewRetranslator(cfg Config) Retranslator {
	events := make(chan apartment.Event, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	ctx, cancel := context.WithCancel(cfg.Ctx)

	consumer := consumer.NewDbConsumer(
		ctx,
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		cfg.Repo,
		events)
	producer := producer.NewKafkaProducer(
		ctx,
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
		ctx: 		ctx,
		cancelFunc: cancel,
	}
}

func (r *retranslator) Start() {
	r.producer.Start()
	r.consumer.Start()
}

func (r *retranslator) Close() {
	r.cancelFunc()
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.StopWait()
}
