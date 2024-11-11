package model

import (
	"fmt"
)

type Ticket struct {
	Title, Author string
	ID, Year      int
}

func NewTicket(ID, year int, title, author string) Ticket {
	return Ticket{
		ID:     ID,
		Title:  title,
		Author: author,
		Year:   year,
	}
}

func (c *Ticket) String() string {
	return fmt.Sprintf("ID: %v, Title: %s, Author: %v, Year: %v", c.ID, c.Title, c.Author, c.Year)
}
