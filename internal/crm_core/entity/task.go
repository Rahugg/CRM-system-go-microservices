package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name             string    `json:"name" gorm:"varchar(255);not null"`
	Description      string    `gorm:"varchar(255);not null" json:"description"`
	DueDate          time.Time `json:"due_date"`
	AssignedTo       uuid.UUID `json:"assigned_to"`
	AssociatedDealID uint      `json:"associated_deal_id"`
	State            string    `json:"state" gorm:"not null"`
	Votes            []Vote    `json:"votes"`
}
type TaskResult struct {
	Task      Task `json:"task"`
	VoteCount int  `json:"vote_count"`
	UserVoted bool `json:"user_voted"`
}

type TaskInput struct {
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	DueDate          time.Time `json:"due_date"`
	AssignedTo       uuid.UUID `json:"assigned_to"`
	AssociatedDealID uint      `json:"associated_deal_id"`
	State            string    `json:"state"`
}

type TaskEditInput struct {
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	DueDate          time.Time `json:"due_date"`
	AssignedTo       uuid.UUID `json:"assigned_to"`
	AssociatedDealID uint      `json:"associated_deal_id"`
	State            string    `json:"state"`
}

type TaskChanges struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ManagerID    uuid.UUID `json:"manager_id" gorm:"not null"`
	TaskID       uint      `json:"task_id" gorm:"not null"`
	ChangedField string    `json:"changed_field"`
	OldValue     string    `json:"old_value"`
	NewValue     string    `json:"new_value"`
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
