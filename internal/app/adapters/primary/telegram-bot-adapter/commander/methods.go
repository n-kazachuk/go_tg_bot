package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/app/adapters/primary/telegram-bot-adapter/path"
	"log"
)

func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath *path.CallbackPath) {
	switch callbackPath.CallbackName {
	default:
		log.Printf("Commander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *Commander) HandleCommand(msg *tgbotapi.Message, commandPath *path.CommandPath) {
	switch commandPath.CommandName {
	case StartCommand:
		c.Start(msg)
	case HelpCommand:
		c.Help(msg)
	case FindCommand:
		c.Find(msg)
	case StopCommand:
		c.Stop(msg)
	default:
		c.Default(msg)
	}
}

func (c *Commander) GetAvailableCommands() []string {
	return []string{
		StartCommand,
		HelpCommand,
		FindCommand,
		StopCommand,
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
