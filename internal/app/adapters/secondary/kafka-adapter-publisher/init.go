package kafka_adapter_publisher

import (
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/n-kazachuk/go_tg_bot/internal/app/config"
	"strings"
)

const (
	ticketFindRequestTopic = "ticket_find_request"
)

var errUnknownType = errors.New("unknown event type")

type KafkaAdapterPublisher struct {
	cfg *config.KafkaConfig

	producer *kafka.Producer
}

func New(cfg *config.KafkaConfig) *KafkaAdapterPublisher {
	kfkCfg := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(cfg.Brokers, ","),
	}

	producer, err := kafka.NewProducer(kfkCfg)
	if err != nil {
		panic(err)
	}

	return &KafkaAdapterPublisher{cfg, producer}
}
