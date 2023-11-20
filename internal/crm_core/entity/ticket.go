package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusTicket string

const (
	Open             StatusTicket = "OPEN"
	InProgressTicket StatusTicket = "IN-PROGRESS"
	Closed           StatusTicket = "CLOSED"
)

type Ticket struct {
	gorm.Model
	IssueDescription string       `gorm:"varchar(255);not null" json:"issue_description"`
	Status           StatusTicket `gorm:"status_ticket"`
	ContactID        uint         `json:"contact_id"`
	AssignedTo       uuid.UUID    `json:"assigned_to"`
}
