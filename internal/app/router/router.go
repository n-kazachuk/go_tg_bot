package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/app/commander"
	"github.com/n-kazachuk/go_tg_bot/internal/app/path"
	"log"
	"runtime/debug"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath path.CommandPath)
}

type Router struct {
	bot       *tgbotapi.BotAPI
	commander Commander
}

func NewRouter(
	bot *tgbotapi.BotAPI,
) *Router {
	return &Router{
		bot:       bot,
		commander: commander.NewCommander(bot),
	}
}

// HandleUpdate used to rote request to his handlers
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
	if !msg.IsCommand() {
		r.showCommandFormat(msg)

		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", msg.Command(), err)
		return
	}

	r.commander.HandleCommand(msg, commandPath)
}

func (r *Router) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}")

	_, err := r.bot.Send(outputMsg)
	if err != nil {
		log.Printf("Router.showCommandFormat: error sending reply message to chat - %v", err)
	}
}
