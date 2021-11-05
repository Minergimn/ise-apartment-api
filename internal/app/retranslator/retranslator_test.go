package retranslator

import (
	"errors"
	"fmt"
	apartment "github.com/ozonmp/ise-apartment-api/internal/model"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/ise-apartment-api/internal/mocks"
)

func TestStart(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)
	cfg := getConfig(repo, sender)

	repo.EXPECT().Lock(gomock.Any()).AnyTimes()

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
	cfg := getConfig(repo, sender)

	eventCount := 10
	events := generateEvents(eventCount)
	lockedEvents := make(map[uint64]bool)
	sendedEvents := make(map[uint64]apartment.ApartmentEvent)
	var eventsLock sync.Mutex
	var sendedEventsLock sync.Mutex

	mockMethodRepoLock(repo, eventsLock, events, lockedEvents)

	repo.EXPECT().Remove(gomock.Any()).DoAndReturn(func(ids []uint64) (err error) {
		eventsLock.Lock()
		defer eventsLock.Unlock()

		for _, id := range ids {
			events[id].Status = apartment.Processed
		}
		return nil
	}).AnyTimes()

	sender.EXPECT().Send(gomock.Any()).DoAndReturn(func(e *apartment.ApartmentEvent) (err error) {
		sendedEventsLock.Lock()
		defer sendedEventsLock.Unlock()

		if _, ok := sendedEvents[e.ID]; ok {
			return nil
		} else {
			sendedEvents[e.ID] = *e
		}

		return nil
	}).AnyTimes()

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	time.Sleep(time.Second * 10)
	retranslator.Close()

	if len(sendedEvents) != eventCount {
		t.Errorf("Not all events were sended to kafka %d/%d", len(sendedEvents), eventCount)
		t.FailNow()
	}

	for _, event := range events {
		if event.Status != apartment.Processed {
			t.Errorf("Some event not removed from queue")
			t.FailNow()
		}
	}
}

func TestSendingFail_AllEventMustBeUnlocked(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)
	cfg := getConfig(repo, sender)

	eventCount := 10
	events := generateEvents(eventCount)
	lockedEvents := make(map[uint64]bool)
	var eventsLock sync.Mutex

	mockMethodRepoLock(repo, eventsLock, events, lockedEvents)

	sender.EXPECT().Send(gomock.Any()).DoAndReturn(func(e *apartment.ApartmentEvent) (err error) {
		return errors.New(fmt.Sprintf("Fail to send event #%d", e.ID))
	}).AnyTimes()
	repo.EXPECT().Unlock(gomock.Any()).DoAndReturn(func(ids []uint64) (err error) {
		eventsLock.Lock()
		defer eventsLock.Unlock()

		for _, id := range ids {
			lockedEvents[id] = false
		}

		return nil
	}).AnyTimes()

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	time.Sleep(time.Second * 10)
	retranslator.Close()

	for _, status := range lockedEvents {
		if status {
			t.Errorf("Some event is locked")
			t.FailNow()
		}
	}
}

func getConfig(repo *mocks.MockEventRepo, sender *mocks.MockEventSender) Config {
	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  5,
		ConsumeSize:    2,
		ConsumeTimeout: 10 * time.Second,
		ProducerCount:  10,
		WorkerCount:    10,
		Repo:           repo,
		Sender:         sender,
	}
	return cfg
}

func generateEvents(eventCount int) []apartment.ApartmentEvent {
	events := make([]apartment.ApartmentEvent, 0)
	for i := uint64(0); i < uint64(eventCount); i++ {
		events = append(events, apartment.ApartmentEvent{
			ID:     i,
			Type:   apartment.Created,
			Status: apartment.Deferred,
			Entity: &apartment.Apartment{
				ID:     i,
				Object: fmt.Sprintf("Object %d", i),
				Owner:  fmt.Sprintf("Owner of object %d", i),
			},
		})
	}
	return events
}

func mockMethodRepoLock(repo *mocks.MockEventRepo, eventsLock sync.Mutex, events []apartment.ApartmentEvent, lockedEvents map[uint64]bool) *gomock.Call {
	return repo.EXPECT().Lock(gomock.Any()).DoAndReturn(func(count uint64) (result []apartment.ApartmentEvent, err error) {
		eventsLock.Lock()
		defer eventsLock.Unlock()

		counter := uint64(0)
		for _, event := range events {
			if _, ok := lockedEvents[event.ID]; ok {
				continue
			} else {
				result = append(result, event)

				lockedEvents[event.ID] = true

				counter++
				if counter == count {
					return
				}
			}
		}

		return nil, nil
	}).AnyTimes()
}
