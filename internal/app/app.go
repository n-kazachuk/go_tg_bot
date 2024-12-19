package app

import (
	telegramBotAdapter "github.com/n-kazachuk/go_tg_bot/internal/app/adapters/primary/telegram-bot-adapter"
	kafkaAdapterPublisher "github.com/n-kazachuk/go_tg_bot/internal/app/adapters/secondary/kafka-adapter-publisher"
	userContextRepositoryDummy "github.com/n-kazachuk/go_tg_bot/internal/app/adapters/secondary/repositories/user-context-repository-dummy"
	"github.com/n-kazachuk/go_tg_bot/internal/app/config"

	"github.com/n-kazachuk/go_tg_bot/internal/app/application/usecases"
	"log/slog"
)

type App struct {
	TelegramBotAdapter *telegramBotAdapter.TelegramBotAdapter
}

func New(log *slog.Logger, cfg *config.Config) *App {
	userContextRepository := userContextRepositoryDummy.NewUserContextRepository()
	kfkAdapterPublisher := kafkaAdapterPublisher.New(&cfg.Kafka)

	usc := usecases.New(
		log,
		userContextRepository,
		kfkAdapterPublisher,
	)

	botAdapter := telegramBotAdapter.New(log, &cfg.Telegram, usc)

	return &App{
		botAdapter,
	}
}
