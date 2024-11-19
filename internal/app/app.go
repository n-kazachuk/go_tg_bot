package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/app/router"
	"github.com/n-kazachuk/go_tg_bot/internal/services"
	"github.com/n-kazachuk/go_tg_bot/internal/services/kafka"
	"github.com/n-kazachuk/go_tg_bot/internal/storages"
	"github.com/n-kazachuk/go_tg_bot/internal/storages/dummy"
	"log"
	"os"
)

type App struct {
	Api    *tgbotapi.BotAPI
	Router *router.Router
}

func New() *App {
	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Panic("environment variable TOKEN not found in .env")
	}

	telegramBotApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	kfkService := kafka.New()
	service := services.New(kfkService)

	contextStorage := dummy.NewContextStorage()
	strg := storages.New(contextStorage)

	routerHandler := router.New(telegramBotApi, service, strg)

	return &App{
		Api:    telegramBotApi,
		Router: routerHandler,
	}
}

// MustRun runs parser and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	cmdCfg := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "find",
			Description: "Поиск билета",
		},
		tgbotapi.BotCommand{
			Command:     "stop",
			Description: "Остановить поиск",
		},
		tgbotapi.BotCommand{
			Command:     "help",
			Description: "Помощь",
		},
	)

	_, err := a.Api.Request(cmdCfg)
	if err != nil {
		return err
	}

	// Uncomment if you want debugging
	a.Api.Debug = true

	log.Printf("Authorized on account %s", a.Api.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := a.Api.GetUpdatesChan(u)

	for update := range updates {
		a.Router.HandleUpdate(update)
	}

	return nil
}

func (a *App) Stop() {
	log.Printf("App stopped")
}
