package commander

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/domain/model"
	"log"
)

const (
	InitFindStep = iota
	FromCityFindStep
	ToCityFindStep
	DateFindStep
	FromTimeFindStep
	ToTimeFindStep
)

var userTicketRequests = make(map[int64]*model.TicketRequest)

func (c *Commander) Find(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest, exists := userTicketRequests[userId]

	if !exists {
		ticketRequest = model.NewTicketRequest()
		ticketRequest.Step = InitFindStep

		userTicketRequests[userId] = ticketRequest
	}

	step := ticketRequest.Step
	text := inputMessage.Text

	switch step {
	case InitFindStep:
		userContext, err := c.storages.ContextStorage.GetContext(userId)
		if err != nil {
			log.Printf("Find %v step: error getting context - %v", InitFindStep, err)
		}

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Укажите откуда (формат: Минск):"),
		)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("Find %v step: error sending reply message to chat - %v", InitFindStep, err)
		}

		userContext.ActiveCommand = FindCommand

		ticketRequest.Step++
		break
	case FromCityFindStep:
		ticketRequest.FromCity = text

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Укажите куда (формат: Москва):"),
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Find %v step: error sending reply message to chat - %v", FromCityFindStep, err)
		}

		ticketRequest.Step++
		break
	case ToCityFindStep:
		ticketRequest.ToCity = text

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Укажите дату (формат: 01.01.2024):"),
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Find %v step: error sending reply message to chat - %v", ToCityFindStep, err)
		}

		ticketRequest.Step++
		break
	case DateFindStep:
		ticketRequest.Date = text

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Укажите с какого времени (формат: 13:00):"),
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Find %v step: error sending reply message to chat - %v", DateFindStep, err)
		}

		ticketRequest.Step++
		break
	case FromTimeFindStep:
		ticketRequest.FromTime = text

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Укажите по какое время (формат: 18:00):"),
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Find %v step: error sending reply message to chat - %v", FromTimeFindStep, err)
		}

		ticketRequest.Step++
		break
	case ToTimeFindStep:
		ticketRequest.ToTime = text

		err := c.services.RequestService.Add(ticketRequest)
		if err != nil {
			log.Printf("Find %v step: error sending reply message to chat - %v", ToTimeFindStep, err)
		}

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf(
				"Начинаю поиск билетов по маршруту %s - %s на дату %s с %s по %s...",
				ticketRequest.FromCity, ticketRequest.ToCity, ticketRequest.Date, ticketRequest.FromTime, ticketRequest.ToTime),
		)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("Find %v step: error sending reply message to chat - %v", ToTimeFindStep, err)
		}

		err = c.storages.ContextStorage.ClearContext(userId)
		if err != nil {
			log.Printf("Find %v step: error sending reply message to chat - %v", ToTimeFindStep, err)
		}

		delete(userTicketRequests, userId)
		break
	}
}
