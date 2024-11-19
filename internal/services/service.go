package services

import (
	model2 "github.com/n-kazachuk/go_tg_bot/internal/domain/model"
)

type RequestService interface {
	Add(ticket *model2.TicketRequest) error
	Remove(ticket *model2.Ticket) error
}

type Service struct {
	RequestService RequestService
}

func New(requestService RequestService) *Service {
	return &Service{
		RequestService: requestService,
	}
}
