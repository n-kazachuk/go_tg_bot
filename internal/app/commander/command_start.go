package commander

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Start(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Привет, %s 👋 \n"+
			"Это бот по поиску билетов. Нажми \"Поиски\" 🔍, чтобы начать", inputMessage.From.FirstName),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Start: error sending reply message to chat - %v", err)
	}
}
