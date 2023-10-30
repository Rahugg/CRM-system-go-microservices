package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Description      string    `gorm:"varchar(255);not null" json:"description"`
	DueDate          time.Time `json:"due_date"`
	AssignedTo       uuid.UUID `json:"assigned_to"`
	AssociatedDealID uint      `json:"associated_deal_id"`
}

//sample JSON
/*
{
    "description": "Sample task description",
    "due_date": "2023-10-30T12:00:00Z",
    "assigned_to": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
    "associated_deal_id": 123456
}
*/
