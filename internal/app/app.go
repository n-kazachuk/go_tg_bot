package app

import (
	osSignalAdapter "github.com/n-kazachuk/go_tg_bot/internal/app/adapters/primary/os-signal-adapter"
	telegramBotAdapter "github.com/n-kazachuk/go_tg_bot/internal/app/adapters/primary/telegram-bot-adapter"
	kafkaAdapterPublisher "github.com/n-kazachuk/go_tg_bot/internal/app/adapters/secondary/kafka-adapter-publisher"
	userContextRepositoryDummy "github.com/n-kazachuk/go_tg_bot/internal/app/adapters/secondary/repositories/user-context-repository-dummy"

	"github.com/n-kazachuk/go_tg_bot/internal/app/application/usecases"
	"github.com/n-kazachuk/go_tg_bot/internal/app/config"
	"log/slog"
)

type App struct {
	OsSignalAdapter    *osSignalAdapter.OsSignalAdapter
	TelegramBotAdapter *telegramBotAdapter.TelegramBotAdapter
}

func New(log *slog.Logger, cfg *config.Config) *App {
	userContextRepository := userContextRepositoryDummy.New()
	kfkAdapterPublisher := kafkaAdapterPublisher.New(&cfg.Kafka)

	usc := usecases.New(
		log,
		userContextRepository,
		kfkAdapterPublisher,
	)

	osAdapter := osSignalAdapter.New(log)
	botAdapter := telegramBotAdapter.New(log, &cfg.Telegram, usc)

	return &App{
		osAdapter,
		botAdapter,
	}
}
