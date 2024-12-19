package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/app/application/usecases"
)

const (
	StartCommand = "start"
	HelpCommand  = "help"
	FindCommand  = "1 🔍"
	StopCommand  = "2 ❎"

	DefaultCommand = "default"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	service *usecases.UseCases
}

func NewCommander(bot *tgbotapi.BotAPI, service *usecases.UseCases) *Commander {
	return &Commander{bot, service}
}
