package main

import (
	osSignalAdapter "github.com/n-kazachuk/go_tg_bot/internal/app/adapters/primary/os-signal-adapter"
	"github.com/n-kazachuk/go_tg_bot/internal/app/config"

	"context"
	"github.com/n-kazachuk/go_tg_bot/internal/app"
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
		graceful.NewProcess(osSignalAdapter.New(log /*, cancel*/)),
	)

	gr.SetLogger(log)
	gr.Start(ctx)
}
