package kafka_adapter_publisher

import "github.com/n-kazachuk/go_tg_bot/internal/app/config"

type KafkaAdapterPublisher struct {
	cfg *config.KafkaConfig
}

func New(cfg *config.KafkaConfig) *KafkaAdapterPublisher {
	return &KafkaAdapterPublisher{cfg}
}
