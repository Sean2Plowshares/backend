package models

import "time"

// TicketType represents the type of ticket.
type TicketType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Ticket represents a ticket
type Ticket struct {
	ID          int64      `json:"id"`
	CreatedDate time.Time  `json:"created_date"`
	UpdatedDate time.Time  `json:"updated_date"`
	Key         string     `json:"key"`
	Summary     string     `json:"summary"`
	Description string     `json:"description"`
	Fields      []Field    `json:"fields"`
	Type        TicketType `json:"ticket_type"`
	Reporter    User       `json:"reporter"`
	Assignee    User       `json:"assignee"`
	Status      Status     `json:"status"`
}

// Status represents a ticket's current status.
type Status struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
