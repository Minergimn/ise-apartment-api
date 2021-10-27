package retranslator

import (
	"fmt"
	apartment "github.com/ozonmp/omp-demo-api/internal/model"
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

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	events := make([]apartment.ApartmentEvent, 0)
	for i := uint64(0); i < 10; i++ {
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
	sendedEvents := make([]apartment.ApartmentEvent, 0)

	repo.EXPECT().Lock(gomock.Any()).DoAndReturn(func(count uint64) ([]apartment.ApartmentEvent, error){
		return events, nil
	}).Times(1)
	repo.EXPECT().Remove(gomock.Any()).DoAndReturn(func(ids []uint64) (err error){
		for _, id := range ids{
			events[id].Status = apartment.Processed
		}
		return nil
	}).AnyTimes()

	sender.EXPECT().Send(gomock.Any()).DoAndReturn(func(e *apartment.ApartmentEvent) (err error){
		sendedEvents = append(sendedEvents, *e)
		return nil
	}).AnyTimes()

	cfg := Config{
		ChannelSize:   512,
		ConsumerCount: 1,
		ConsumeSize:   10,
		ConsumeTimeout: 10 * time.Second,
		ProducerCount: 10,
		WorkerCount:   10,
		Repo:          repo,
		Sender:        sender,
	}

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	time.Sleep(time.Second * 15)
	retranslator.Close()

	if len(sendedEvents) < len(events){
		t.Errorf("Not all events were sended to kafka %d/%d", len(sendedEvents), len(events))
		t.FailNow()
	}

	for _, event := range events{
		if event.Status != apartment.Processed {
			t.Errorf("Some event not removed from queue")
			t.FailNow()
		}
	}
}
