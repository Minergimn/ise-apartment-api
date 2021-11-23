package producer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ozonmp/ise-apartment-api/internal/metrics"
	apartment "github.com/ozonmp/ise-apartment-api/internal/model"

	"github.com/ozonmp/ise-apartment-api/internal/logger"

	"github.com/ozonmp/ise-apartment-api/internal/app/repo"
	"github.com/ozonmp/ise-apartment-api/internal/app/sender"

	"github.com/gammazero/workerpool"
)

//Producer comment for linter
type Producer interface {
	Start(context.Context)
	Close()
}

type producer struct {
	n       uint64
	timeout time.Duration //nolint

	sender sender.EventSender
	events <-chan apartment.Event

	repo repo.EventRepo

	workerPool *workerpool.WorkerPool

	wg   *sync.WaitGroup
}

//NewKafkaProducer comment for linter
func NewKafkaProducer(
	n uint64,
	sender sender.EventSender,
	events <-chan apartment.Event,
	repo repo.EventRepo,
	workerPool *workerpool.WorkerPool,
) Producer {

	wg := &sync.WaitGroup{}

	return &producer{
		n:          n,
		sender:     sender,
		events:     events,
		repo:       repo,
		workerPool: workerPool,
		wg:         wg,
	}
}

func (p *producer) Start(ctx context.Context) {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					if err := p.sender.Send(&event); err != nil {
						p.workerPool.Submit(func() {
							logger.DebugKV(ctx, fmt.Sprintf("Sending of event %s failed with error: %t", event.String(), err))

							if err = p.repo.Unlock(ctx, []uint64{event.ID}); err != nil {
								logger.DebugKV(ctx, fmt.Sprintf("Unlock event %s has error: %t", event.String(), err))
							}
						})
					} else {
						p.workerPool.Submit(func() {
							logger.DebugKV(ctx, fmt.Sprintf("Sending of event %d was succeeded", event.ID))

							if err = p.repo.Remove(ctx, []uint64{event.ID}); err != nil {
								logger.DebugKV(ctx, fmt.Sprintf("Removing event %s has error: %t", event.String(), err))
							}

							metrics.SubCurrentRetranslatorEventsCount(1)
						})
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}

func (p *producer) Close() {
	p.wg.Wait()
}
