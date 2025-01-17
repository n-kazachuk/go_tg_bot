package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"ℹ️ Помощь \n\n"+
			"Этот бот занимается поиском билетов на сторонних ресурсах. \n"+
			"Для того чтобы начать поиск, выбери '"+FindCommand+"' в нижнем меню. \n"+
			"Далее заполни параметры поиска, после бот будет искать билеты и пришлет уведомление, если что-то найдет. \n\n"+
			FindCommand+" - начать новый поиск \n"+
			ListCommand+" - список твоих активных поисков \n"+
			StopCommand+" - остановить поиск \n",
	)

	c.Send(msg)
	c.Default(inputMessage)
}
