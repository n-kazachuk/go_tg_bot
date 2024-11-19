package kafka

import (
	model2 "github.com/n-kazachuk/go_tg_bot/internal/domain/model"
	"log"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Add(ticket *model2.TicketRequest) error {
	log.Print("ticket sent to kafka")
	return nil
}

func (s *Service) Remove(ticket *model2.Ticket) error {
	log.Print("ticket sent to kafka")
	return nil
}
