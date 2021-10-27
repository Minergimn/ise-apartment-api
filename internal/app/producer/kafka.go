package producer

import (
	"log"
	"sync"
	"time"

	"github.com/ozonmp/omp-demo-api/internal/app/repo"
	"github.com/ozonmp/omp-demo-api/internal/app/sender"
	"github.com/ozonmp/omp-demo-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Producer interface {
	Start()
	Close()
}

type producer struct {
	n       uint64
	timeout time.Duration

	sender sender.EventSender
	events <-chan apartment.ApartmentEvent

	repo repo.EventRepo

	workerPool *workerpool.WorkerPool

	wg   *sync.WaitGroup
	done chan bool
}

func NewKafkaProducer(
	n uint64,
	sender sender.EventSender,
	events <-chan apartment.ApartmentEvent,
	repo repo.EventRepo,
	workerPool *workerpool.WorkerPool,
) Producer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &producer{
		n:          n,
		sender:     sender,
		events:     events,
		repo: 		repo,
		workerPool: workerPool,
		wg:         wg,
		done:       done,
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					if err := p.sender.Send(&event); err != nil {
						p.workerPool.Submit(func() {
							log.Printf("Sending of event %s failed with error: %t", event.String(), err)

							if err = p.repo.Unlock([]uint64{event.ID}); err != nil{
								log.Printf("Unlock event %s has error: %t",event.String(), err)
							}
						})
					} else {
						p.workerPool.Submit(func() {
							log.Printf("Sending of event %d was succeeded", event.ID)

							if err = p.repo.Remove([]uint64{event.ID}); err != nil{
								log.Printf("Removing event %s has error: %t",event.String(), err)
							}
						})
					}
				case <-p.done:
					return
				}
			}
		}()
	}
}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}
