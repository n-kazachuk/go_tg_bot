package commander

import (
	ticketsRequest "github.com/n-kazachuk/go_tg_bot/internal/app/domain/tickets-request"

	"errors"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/n-kazachuk/go_tg_bot/internal/libs/logger/sl"
	"regexp"
	"strings"
	"time"
)

const (
	DateFormat = "02.01.2006"
	TimeFormat = "15:04"

	InitFindStep = iota
	FromCityFindStep
	ToCityFindStep
	DateFindStep
	FromTimeFindStep
	ToTimeFindStep
)

var userTicketsRequests = make(map[int64]*ticketsRequest.TicketsRequest)

func (c *Commander) Find(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest, exists := userTicketsRequests[userId]

	if !exists {
		ticketRequest = ticketsRequest.New()
		ticketRequest.Step = InitFindStep

		userTicketsRequests[userId] = ticketRequest
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
	default:
		c.Default(inputMessage)
	}
}

func (c *Commander) askFromCity(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Укажите откуда (формат: Минск):",
	)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	c.Send(msg)

	userContext, _ := c.service.GetUserContext(userId)
	userContext.ActiveCommand = FindCommand

	c.updateStep(inputMessage)
}

func (c *Commander) handleFromCity(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketsRequests[userId]

	if err := c.validateCity(inputMessage.Text); err != nil {
		c.SendError(inputMessage, err)
		return
	}

	ticketRequest.FromCity = strings.Title(inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Укажите куда (формат: Москва):",
	)

	c.Send(msg)
	c.updateStep(inputMessage)
}

func (c *Commander) handleToCity(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketsRequests[userId]

	if err := c.validateCity(inputMessage.Text); err != nil {
		c.SendError(inputMessage, err)
		return
	}

	ticketRequest.ToCity = strings.Title(inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Укажите дату (формат: 01.01.2024):",
	)

	c.Send(msg)
	c.updateStep(inputMessage)
}

func (c *Commander) handleDate(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketsRequests[userId]

	if err := c.validateDate(inputMessage.Text); err != nil {
		c.SendError(inputMessage, err)
		return
	}

	ticketRequest.Date, _ = time.Parse(DateFormat, inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Укажите с какого времени (формат: 13:00):",
	)

	c.Send(msg)
	c.updateStep(inputMessage)
}

func (c *Commander) handleFromTime(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketsRequests[userId]

	if err := c.validateTime(inputMessage.Text); err != nil {
		c.SendError(inputMessage, err)
		return
	}

	ticketRequest.FromTime, _ = time.Parse(TimeFormat, inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Укажите по какое время (формат: 18:00):",
	)

	c.Send(msg)
	c.updateStep(inputMessage)
}

func (c *Commander) handleToTime(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketsRequests[userId]

	if err := c.validateTime(inputMessage.Text); err != nil {
		c.SendError(inputMessage, err)
		return
	}

	ticketRequest.ToTime, _ = time.Parse(TimeFormat, inputMessage.Text)

	err := c.service.SendTicketSearchRequest(ticketRequest)
	if err != nil {
		c.log.Error("Error adding ticket request", sl.Err(err))
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf(
			"Начинаю поиск билетов (*%s*) 🕵: \n"+
				"*Из:* %s \n"+
				"*В:* %s \n"+
				"*На время:* %s - %s",
			ticketRequest.Date.Format(DateFormat), ticketRequest.FromCity, ticketRequest.ToCity,
			ticketRequest.FromTime.Format(TimeFormat), ticketRequest.ToTime.Format(TimeFormat)),
	)

	msg.ParseMode = tgbotapi.ModeMarkdown

	c.Send(msg)

	err = c.service.ClearUserContext(userId)
	if err != nil {
		c.log.Error("Error clearing context", sl.Err(err))
	}

	delete(userTicketsRequests, userId)

	c.Default(inputMessage)
}

func (c *Commander) updateStep(inputMessage *tgbotapi.Message) {
	userId := inputMessage.From.ID
	ticketRequest := userTicketsRequests[userId]
	ticketRequest.Step++
}

func (c *Commander) validateCity(city string) error {
	if city == "" {
		return errors.New("передана пустая строка")
	}

	re := regexp.MustCompile("^[А-ЯЁа-яё]+$")
	if !re.MatchString(city) {
		return fmt.Errorf("город указан в неверном формате")
	}

	return nil
}

func (c *Commander) validateDate(date string) error {
	parsedDate, err := time.Parse(DateFormat, date)
	if err != nil {
		return errors.New("неверный формат даты")
	}

	currentDate := time.Now().Truncate(24 * time.Hour)

	if parsedDate.Before(currentDate) {
		return errors.New("дата не может быть меньше текущей")
	}

	return nil
}

func (c *Commander) validateTime(t string) error {
	_, err := time.Parse(TimeFormat, t)
	if err != nil {
		return errors.New("неверный формат времени")
	}

	return nil
}
