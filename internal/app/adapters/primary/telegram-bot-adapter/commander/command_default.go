package commander

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"1. Начать поиск \n"+
			"2. Остановить поиск",
	)

	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("1 🔍"),
			tgbotapi.NewKeyboardButton("2 ❎"),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Default: error sending reply message to chat - %v", err)
	}
}
