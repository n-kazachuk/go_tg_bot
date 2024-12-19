package usecases

import (
	"github.com/n-kazachuk/go_tg_bot/internal/app/domain/model"
	"log/slog"
)

type UseCases struct {
	log                   *slog.Logger
	userContextRepository userContextRepository
	tickerRequestSender   tickerRequestSender
}

type userContextRepository interface {
	GetContext(userId int64) (*model.UserContext, error)
	ClearContext(userID int64) error
}

type tickerRequestSender interface {
	SendTicketRequest(ticket *model.TicketRequest) error
}

func New(
	log *slog.Logger,
	userContextRepository userContextRepository,
	tickerRequestSender tickerRequestSender,
) *UseCases {
	return &UseCases{
		log,
		userContextRepository,
		tickerRequestSender,
	}
}
