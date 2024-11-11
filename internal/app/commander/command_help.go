package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"ℹ️ List of commands: \n\n"+
			"1. /help - Gives you list of all commands \n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Help: error sending reply message to chat - %v", err)
	}
}
