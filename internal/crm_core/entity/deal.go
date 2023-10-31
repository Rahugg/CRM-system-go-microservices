package entity

import (
	"gorm.io/gorm"
)

type StatusDeal string

const (
	Initiated  StatusDeal = "INITIATED"
	InProgress StatusDeal = "IN-PROGRESS"
	ClosedWon  StatusDeal = "CLOSED-WON"
	ClosedLost StatusDeal = "CLOSED-LOST"
)

type Deal struct {
	gorm.Model
	Title     string     `gorm:"varchar(255);not null" json:"title"`
	Value     uint       `gorm:"varchar(255);not null" json:"value"`
	Status    StatusDeal `gorm:"type:status_deal"`
	ContactID uint       `json:"contact_id"`
	RepID     uint       `json:"rep_id"`
}
