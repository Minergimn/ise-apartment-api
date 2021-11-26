package sender

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/ozonmp/ise-apartment-api/internal/model"
)

//EventSender comment for linter
type EventSender interface {
	Send(event *apartment.Event, topic string) error
}

type eventSender struct {
	producer sarama.SyncProducer
}

//NewEventSender comment for linter
func NewEventSender(brokers []string) (EventSender, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return &eventSender{
		producer: producer,
	}, err
}

func (e *eventSender) Send(event *apartment.Event, topic string) error {
	json, err := json.Marshal(event)
	if err == nil {
		err = sendMessage(e.producer, topic, json)
	}

	return err
}

func sendMessage(producer sarama.SyncProducer, topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(message),
	}
	_, _, err := producer.SendMessage(msg)
	return err
}
