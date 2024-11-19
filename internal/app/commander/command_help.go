package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"ℹ️ Список команд: \n\n"+
			"1. /help - Вывести список комманд \n"+
			"1. /find - Поиск билетов \n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Help: error sending reply message to chat - %v", err)
	}
}
