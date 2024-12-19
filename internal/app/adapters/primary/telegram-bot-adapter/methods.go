package telegram_bot_adapter

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/libs/helpers"
)

func (a *TelegramBotAdapter) Start(ctx context.Context) error {
	cmdCfg := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "start",
			Description: "Начать",
		},
		tgbotapi.BotCommand{
			Command:     "help",
			Description: "Помощь",
		},
	)

	_, err := a.api.Request(cmdCfg)
	if err != nil {
		return err
	}

	a.log.Info(fmt.Sprintf("%s: Authorized on account %s", helpers.GetFunctionName(), a.api.Self.UserName))

	u := tgbotapi.UpdateConfig{
		Timeout: a.cfg.Timeout,
	}

	updates := a.api.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case update := <-updates:
			a.router.HandleUpdate(update)
		}
	}
}
