package commander

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "❌ Undefined command: "+inputMessage.Text+". Write /help for help")

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Help: error sending reply message to chat - %v", err)
	}
}
