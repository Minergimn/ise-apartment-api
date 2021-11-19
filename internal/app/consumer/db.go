package consumer

import (
	"context"
	"fmt"
	"github.com/ozonmp/ise-apartment-api/internal/logger"
	"github.com/ozonmp/ise-apartment-api/internal/metrics"
	"sync"
	"time"

	"github.com/ozonmp/ise-apartment-api/internal/app/repo"
	"github.com/ozonmp/ise-apartment-api/internal/model"
)

//Consumer comment for linter
type Consumer interface {
	Start(ctx context.Context)
	Close(ctx context.Context)
}

type consumer struct {
	n      uint64
	events chan<- apartment.ApartmentEvent

	repo repo.EventRepo

	batchSize uint64
	timeout   time.Duration

	cancelFunc context.CancelFunc
	wg   *sync.WaitGroup
}

//Config comment for linter
type Config struct {
	n         uint64 //nolint
	events    chan<- apartment.ApartmentEvent //nolint
	repo      repo.EventRepo //nolint
	batchSize uint64 //nolint
	timeout   time.Duration //nolint
}

//NewDbConsumer comment for linter
func NewDbConsumer(
	n uint64,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo repo.EventRepo,
	events chan<- apartment.ApartmentEvent) Consumer {

	wg := &sync.WaitGroup{}

	return &consumer{
		n:         n,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		repo:      repo,
		events:    events,
		wg:        wg,
	}
}

func (c *consumer) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	c.cancelFunc = cancel

	for i := uint64(0); i < c.n; i++ {
		c.wg.Add(1)

		go func() {
			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)
			for {
				select {
				case <-ticker.C:
					events, err := c.repo.Lock(ctx, c.batchSize)
					if err != nil {
						continue
					}
					for _, event := range events {
						c.events <- event
						metrics.AddCurrentRetranslatorEventsCount(1)
						logger.DebugKV(ctx, fmt.Sprintf("Add event %d to retraslator from db", event.ID))
					}
				case <-ctx.Done():
					ticker.Stop()
					return
				}
			}
		}()
	}
}

func (c *consumer) Close(ctx context.Context) {
	c.cancelFunc()
	c.wg.Wait()
}
