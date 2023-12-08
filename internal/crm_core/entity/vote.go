package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vote struct {
	gorm.Model
	SenderID uuid.UUID `json:"sender_id"`
	TaskID   uint      `json:"task_id"`
	Task     *Task     `gorm:"foreignkey:TaskID" json:"todo"`
}

type VoteInput struct {
	TaskID uint `json:"task_id"`
}
