package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/n-kazachuk/go_tg_bot/internal/app/path"
	"github.com/n-kazachuk/go_tg_bot/internal/services"
	"log"
)

type CommandHandler interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type Commander struct {
	bot           *tgbotapi.BotAPI
	ticketService *services.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
) *Commander {
	ticketService := services.NewService()

	return &Commander{
		bot:           bot,
		ticketService: ticketService,
	}
}

func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	default:
		log.Printf("Commander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *Commander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	default:
		c.Default(msg)
	}
}

func (c *Commander) SendError(inputMessage *tgbotapi.Message, err error) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"❌ Error!!! \n"+
			err.Error(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("Commander.HandleError: error sending reply message to chat - %v", err)
	}
}
