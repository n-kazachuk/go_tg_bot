package main

import (
	"context"
	"github.com/n-kazachuk/go_tg_bot/internal/app"
	"github.com/n-kazachuk/go_tg_bot/internal/app/config"
	"github.com/n-kazachuk/go_tg_bot/internal/libs/graceful"
	"github.com/n-kazachuk/go_tg_bot/internal/libs/slogpretty"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.MustLoad()
	log := slogpretty.SetupLogger(cfg.Env)

	application := app.New(log, cfg)

	gr := graceful.New(
		graceful.NewProcess(application.TelegramBotAdapter),
		graceful.NewProcess(application.OsSignalAdapter),
	)

	gr.SetLogger(log)
	gr.Start(ctx)
}
