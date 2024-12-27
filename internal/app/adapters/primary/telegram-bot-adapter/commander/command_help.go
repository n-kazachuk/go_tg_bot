package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"ℹ️ Список команд: \n\n"+
			"1. /help - Вывести список комманд \n"+
			"1. /find - Поиск билетов \n",
	)

	c.Send(msg)
}
