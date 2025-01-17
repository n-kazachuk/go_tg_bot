package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/app/application/usecases"
	"log/slog"
)

const (
	StartCommand = "start"
	HelpCommand  = "help"
	FindCommand  = "1 🔍"
	ListCommand  = "2 📃"
	StopCommand  = "3 ❎"

	DefaultCommand = "default"
)

type Commander struct {
	log     *slog.Logger
	bot     *tgbotapi.BotAPI
	service *usecases.UseCases
}

func NewCommander(log *slog.Logger, bot *tgbotapi.BotAPI, service *usecases.UseCases) *Commander {
	return &Commander{log, bot, service}
}
