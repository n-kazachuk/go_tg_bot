package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/app/adapters/primary/telegram-bot-adapter/commander"
	"github.com/n-kazachuk/go_tg_bot/internal/app/adapters/primary/telegram-bot-adapter/path"
	"log"
	"runtime/debug"
	"slices"
)

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v\n%v", panicValue, string(debug.Stack()))
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		r.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		r.handleMessage(update.Message)
	}
}

func (r *Router) handleCallback(callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	r.commander.HandleCallback(callback, callbackPath)
}

func (r *Router) handleMessage(msg *tgbotapi.Message) {
	const op = "Router.handleMessage"

	command := r.getCommand(msg)
	commandPath, err := path.ParseCommand(command)

	if err != nil {
		log.Printf("%s: error parsing callback data `%s` - %v", op, msg.Command(), err)
		return
	}

	r.commander.HandleCommand(msg, commandPath)
}

func (r *Router) getCommand(msg *tgbotapi.Message) string {
	var command string

	activeCommand, status := r.getActiveCommand(msg)
	if status {
		command = activeCommand
	} else if msg.IsCommand() {
		command = msg.Command()
	} else {
		command = msg.Text
	}

	if !slices.Contains(r.commander.GetAvailableCommands(), command) {
		command = commander.DefaultCommand
	}

	return command
}

func (r *Router) getActiveCommand(msg *tgbotapi.Message) (string, bool) {
	const op = "Router.getActiveCommand"

	userContext, err := r.service.GetUserContext(msg.From.ID)
	if err != nil {
		log.Printf("%s: error getting active command - %v", op, err)
		return "", false
	}

	activeCommand := userContext.ActiveCommand
	if activeCommand == "" {
		return "", false
	}

	return activeCommand, true
}
