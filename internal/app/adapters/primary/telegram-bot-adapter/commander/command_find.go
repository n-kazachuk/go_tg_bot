package commander

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/app/domain/model"
	"log"
	"regexp"
	"strings"
	"time"
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

	switch step {
	case InitFindStep:
		c.askFromCity(inputMessage)
	case FromCityFindStep:
		c.handleFromCity(inputMessage)
	case ToCityFindStep:
		c.handleToCity(inputMessage)
	case DateFindStep:
		c.handleDate(inputMessage)
	case FromTimeFindStep:
		c.handleFromTime(inputMessage)
	case ToTimeFindStep:
		c.handleToTime(inputMessage)
	}
}

func (c *Commander) askFromCity(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Укажите откуда (формат: Минск):",
	)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message for step %d: %v", InitFindStep, err)
	}

	userContext, err := c.service.GetUserContext(userId)
	userContext.ActiveCommand = FindCommand

	c.updateStep(inputMessage)
}

func (c *Commander) handleFromCity(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketRequests[userId]

	if !c.validateCity(inputMessage.Text) {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Неверный формат города, пожалуйста, укажите город (формат: Минск).",
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Error sending reply message for step %d: %v", FromCityFindStep, err)
		}
		return
	}

	ticketRequest.FromCity = strings.Title(inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Укажите куда (формат: Москва):",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message for step %d: %v", FromCityFindStep, err)
	}

	c.updateStep(inputMessage)
}

func (c *Commander) handleToCity(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketRequests[userId]

	if !c.validateCity(inputMessage.Text) {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Неверный формат города, пожалуйста, укажите город (формат: Москва).",
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Error sending reply message for step %d: %v", ToCityFindStep, err)
		}
		return
	}

	ticketRequest.ToCity = strings.Title(inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Укажите дату (формат: 01.01.2024):",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message for step %d: %v", ToCityFindStep, err)
	}

	c.updateStep(inputMessage)
}

func (c *Commander) handleDate(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketRequests[userId]

	if !c.validateDate(inputMessage.Text) {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Неверный формат даты, пожалуйста, укажите дату (формат: 01.01.2024).",
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Error sending reply message for step %d: %v", DateFindStep, err)
		}
		return
	}

	ticketRequest.Date = inputMessage.Text

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Укажите с какого времени (формат: 13:00):",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message for step %d: %v", DateFindStep, err)
	}

	c.updateStep(inputMessage)
}

func (c *Commander) handleFromTime(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketRequests[userId]

	if !c.validateTime(inputMessage.Text) {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Неверный формат времени, пожалуйста, укажите время (формат: 13:00).",
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Error sending reply message for step %d: %v", FromTimeFindStep, err)
		}
		return
	}

	ticketRequest.FromTime = inputMessage.Text

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Укажите по какое время (формат: 18:00):",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message for step %d: %v", FromTimeFindStep, err)
	}

	c.updateStep(inputMessage)
}

func (c *Commander) handleToTime(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketRequests[userId]

	if !c.validateTime(inputMessage.Text) {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Неверный формат времени, пожалуйста, укажите время (формат: 18:00).",
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Error sending reply message for step %d: %v", ToTimeFindStep, err)
		}
		return
	}

	ticketRequest.ToTime = inputMessage.Text

	err := c.service.SendTicketSearchRequest(ticketRequest)
	if err != nil {
		log.Printf("Error adding ticket request: %v", err)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf(
			"Начинаю поиск билетов по маршруту %s - %s на дату %s с %s по %s...",
			ticketRequest.FromCity, ticketRequest.ToCity, ticketRequest.Date, ticketRequest.FromTime, ticketRequest.ToTime),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message for step %d: %v", ToTimeFindStep, err)
	}

	err = c.service.ClearUserContext(userId)
	if err != nil {
		log.Printf("Error clearing context: %v", err)
	}

	delete(userTicketRequests, userId)

	c.Default(inputMessage)
}

func (c *Commander) updateStep(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketRequests[userId]
	ticketRequest.Step++
}

func (c *Commander) validateCity(city string) bool {
	if city == "" {
		return false
	}

	re := regexp.MustCompile("^[А-ЯЁа-яё]+$")
	if !re.MatchString(city) {
		return false
	}

	return true
}

func (c *Commander) validateDate(date string) bool {
	parsedDate, err := time.Parse("02.01.2006", date)
	if err != nil {
		return false
	}

	currentDate := time.Now().Truncate(24 * time.Hour)

	if parsedDate.Before(currentDate) {
		return false
	}

	return true
}

func (c *Commander) validateTime(t string) bool {
	_, err := time.Parse("15:04", t)
	return err == nil
}
