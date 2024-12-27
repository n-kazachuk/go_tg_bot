package kafka_adapter_publisher

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/n-kazachuk/go_tg_bot/internal/app/domain/tickets-request"
)

// Produce main method to send kafka messages
func (a *KafkaAdapterPublisher) Produce(message []byte, topic string) error {
	kfkMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: message,
		Key:   nil,
	}

	kfkChan := make(chan kafka.Event)

	if err := a.producer.Produce(kfkMsg, kfkChan); err != nil {
		return err
	}

	e := <-kfkChan

	switch ev := e.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return ev
	default:
		return errUnknownType
	}
}

func (a *KafkaAdapterPublisher) SendTicketRequest(ticket *tickets_request.TicketsRequest) error {
	ticketJson, err := json.Marshal(ticket)
	if err != nil {
		return err
	}

	if err := a.Produce(ticketJson, ticketFindRequestTopic); err != nil {
		return err
	}

	return nil
}
