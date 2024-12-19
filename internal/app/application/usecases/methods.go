package usecases

import "github.com/n-kazachuk/go_tg_bot/internal/app/domain/model"

func (s *UseCases) GetUserContext(userId int64) (*model.UserContext, error) {
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

func (s *UseCases) SendTicketSearchRequest(ticket *model.TicketRequest) error {
	err := s.tickerRequestSender.SendTicketRequest(ticket)
	if err != nil {
		return err
	}

	return nil
}
