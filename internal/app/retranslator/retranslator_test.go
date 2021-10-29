package retranslator

import (
	"fmt"
	apartment "github.com/ozonmp/omp-demo-api/internal/model"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/omp-demo-api/internal/mocks"
)

func TestStart(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	repo.EXPECT().Lock(gomock.Any()).AnyTimes()

	cfg := Config{
		ChannelSize:   512,
		ConsumerCount: 2,
		ConsumeSize:   10,
		ConsumeTimeout: 10 * time.Second,
		ProducerCount: 2,
		WorkerCount:   2,
		Repo:          repo,
		Sender:        sender,
	}

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	time.Sleep(time.Second)
	retranslator.Close()
}

func TestSendAndRemove(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)
	
	eventCount := 10

	events := make([]apartment.ApartmentEvent, 0)
	lockedEvents := make(map[uint64]bool)
	sendedEvents := make(map[uint64]apartment.ApartmentEvent)
	var eventsLock sync.Mutex
	var sendedEventsLock sync.Mutex


	for i := uint64(0); i < uint64(eventCount); i++ {
		events = append(events, apartment.ApartmentEvent{
			ID: i,
			Type: apartment.Created,
			Status: apartment.Deferred,
			Entity: &apartment.Apartment{
				ID: i,
				Object: fmt.Sprintf("Object %d", i),
				Owner: fmt.Sprintf("Owner of object %d", i),
			},
		})
	}

	repo.EXPECT().Lock(gomock.Any()).DoAndReturn(func(count uint64) (result []apartment.ApartmentEvent, err error){
		eventsLock.Lock()
		defer eventsLock.Unlock()

		counter := uint64(0)
		for _, event := range events{
			if _, ok := lockedEvents[event.ID]; ok {
				continue
			} else {
				result = append(result, event)

				lockedEvents[event.ID] = true

				counter++
				if counter == count{
					return
				}
			}
		}
		
		return nil, nil
	}).AnyTimes()
	repo.EXPECT().Remove(gomock.Any()).DoAndReturn(func(ids []uint64) (err error){
		eventsLock.Lock()
		defer eventsLock.Unlock()

		for _, id := range ids{
			events[id].Status = apartment.Processed
		}
		return nil
	}).AnyTimes()

	sender.EXPECT().Send(gomock.Any()).DoAndReturn(func(e *apartment.ApartmentEvent) (err error){
		sendedEventsLock.Lock()
		defer sendedEventsLock.Unlock()

		if _, ok := sendedEvents[e.ID]; ok {
			return nil
		} else {
			sendedEvents[e.ID] = *e
		}

		return nil
	}).AnyTimes()

	cfg := Config{
		ChannelSize:   512,
		ConsumerCount: 5,
		ConsumeSize:   2,
		ConsumeTimeout: 10 * time.Second,
		ProducerCount: 10,
		WorkerCount:   10,
		Repo:          repo,
		Sender:        sender,
	}

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	time.Sleep(time.Second * 10)
	retranslator.Close()

	if len(sendedEvents) != eventCount{
		t.Errorf("Not all events were sended to kafka %d/%d", len(sendedEvents), eventCount)
		t.FailNow()
	}

	for _, event := range events{
		if event.Status != apartment.Processed {
			t.Errorf("Some event not removed from queue")
			t.FailNow()
		}
	}
}