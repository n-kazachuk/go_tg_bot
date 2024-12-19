package kafka_adapter_publisher

import (
	"github.com/n-kazachuk/go_tg_bot/internal/app/domain/model"
	"log"
)

func (s *KafkaAdapterPublisher) SendTicketRequest(ticket *model.TicketRequest) error {
	log.Print("ticket sent to kafka")
	return nil
}
