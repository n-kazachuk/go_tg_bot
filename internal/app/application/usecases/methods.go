package usecases

import (
	"github.com/n-kazachuk/go_tg_bot/internal/app/domain/tickets-request"
	"github.com/n-kazachuk/go_tg_bot/internal/app/domain/user-context"
)

func (s *UseCases) GetUserContext(userId int64) (*user_context.UserContext, error) {
	userContext, err := s.userContextRepository.GetContext(userId)
	if err != nil {
		return nil, err
	}

	return userContext, nil
}

func (s *UseCases) ClearUserContext(userId int64) error {
	err := s.userContextRepository.ClearContext(userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UseCases) SendTicketSearchRequest(ticket *tickets_request.TicketsRequest) error {
	err := s.tickerRequestSender.SendTicketRequest(ticket)
	if err != nil {
		return err
	}

	return nil
}
