package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/app/adapters/primary/telegram-bot-adapter/commander"
	"github.com/n-kazachuk/go_tg_bot/internal/app/adapters/primary/telegram-bot-adapter/path"
	"github.com/n-kazachuk/go_tg_bot/internal/app/application/usecases"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath *path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath *path.CommandPath)

	GetAvailableCommands() []string
}

type Router struct {
	bot     *tgbotapi.BotAPI
	service *usecases.UseCases

	commander Commander
}

func New(bot *tgbotapi.BotAPI, service *usecases.UseCases) *Router {
	cmdr := commander.NewCommander(bot, service)

	return &Router{bot, service, cmdr}
}
