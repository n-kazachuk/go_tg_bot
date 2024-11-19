package main

import (
	"github.com/joho/godotenv"
	"github.com/n-kazachuk/go_tg_bot/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_ = godotenv.Load()

	application := app.New()

	go func() {
		application.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
}
