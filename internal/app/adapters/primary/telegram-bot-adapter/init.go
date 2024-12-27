package telegram_bot_adapter

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/app/adapters/primary/telegram-bot-adapter/router"
	"github.com/n-kazachuk/go_tg_bot/internal/app/application/usecases"
	"github.com/n-kazachuk/go_tg_bot/internal/app/config"
	"github.com/n-kazachuk/go_tg_bot/internal/libs/helpers"
	"log/slog"
)

type TelegramBotAdapter struct {
	log *slog.Logger
	cfg *config.TelegramConfig

	api    *tgbotapi.BotAPI
	router *router.Router
}

func New(log *slog.Logger, cfg *config.TelegramConfig, service *usecases.UseCases) *TelegramBotAdapter {
	token := cfg.Token
	if token == "" {
		err := fmt.Errorf("%s: telegram token not found", helpers.GetFunctionName())
		panic(err)
	}

	telegramBotApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		err := fmt.Errorf("%s: can't connect to telegram API", helpers.GetFunctionName())
		panic(err)
	}

	telegramBotApi.Debug = cfg.Debug

	routerHandler := router.New(log, telegramBotApi, service)

	return &TelegramBotAdapter{log, cfg, telegramBotApi, routerHandler}
}
