package os_signal_adapter

import (
	"context"
	"fmt"
	"github.com/n-kazachuk/go_tg_bot/internal/libs/helpers"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type OsSignalAdapter struct {
	log *slog.Logger
	//cancel context.CancelFunc
}

func New(log *slog.Logger /*, cancel context.CancelFunc*/) *OsSignalAdapter {
	return &OsSignalAdapter{log /*, cancel*/}
}

func (a *OsSignalAdapter) Start(ctx context.Context) error {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case sig := <-ch:
		err := fmt.Errorf("%s: system signal getted %s", helpers.GetFunctionName(), sig.String())

		//a.cancel()

		return err
	}
}
