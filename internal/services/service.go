package services

import (
	"fmt"
	"github.com/n-kazachuk/go_tg_bot/internal/model"
)

var entitiesPrimaryKey = 0
var allEntities = make([]model.Ticket, 0)

type TicketService interface {
	Describe(ticketID uint64) (*model.Ticket, error)
	List(cursor uint64, limit uint64) ([]model.Ticket, error)
	Create(ticket model.Ticket) (uint64, error)
	Update(ticketID uint64, ticket model.Ticket) error
	Remove(ticketID uint64) (bool, error)
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List(cursor uint64, limit uint64) ([]model.Ticket, error) {
	if len(allEntities) == 0 {
		return []model.Ticket{}, nil
	}

	from := cursor * limit
	if from > uint64(len(allEntities)) {
		return nil, fmt.Errorf("❌ Invalid data for paginate")
	}

	to := from + limit
	if to > uint64(len(allEntities)) {
		to = uint64(len(allEntities))
	}

	return allEntities[from:to], nil
}

func (s *Service) Describe(ticketID uint64) (*model.Ticket, error) {
	for index, ticket := range allEntities {
		if ticket.ID == int(ticketID) {
			return &allEntities[index], nil
		}
	}

	return nil, fmt.Errorf("❌ Invalid ID")
}

func (s *Service) Create(ticket model.Ticket) (uint64, error) {
	ticket.ID = s.generateID()

	allEntities = append(allEntities, ticket)

	return uint64(ticket.ID), nil
}

func (s *Service) Update(ticketID uint64, ticket model.Ticket) error {
	for index, entity := range allEntities {
		if entity.ID == int(ticketID) {
			allEntities[index] = ticket
			return nil
		}
	}

	return fmt.Errorf("❌ Invalid ID")
}

func (s *Service) Remove(ticketID uint64) (bool, error) {
	for index, ticket := range allEntities {
		if ticket.ID == int(ticketID) {
			allEntities = append(allEntities[:index], allEntities[index+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("❌ Invalid ID")
}

func (s *Service) GetCoursesCount() int {
	return len(allEntities)
}

func (s *Service) generateID() int {
	entitiesPrimaryKey += 1
	return entitiesPrimaryKey
}
