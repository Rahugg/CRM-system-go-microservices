package entity

import (
	"database/sql/driver"
	"gorm.io/gorm"
)

type StatusDeal string

const (
	Initiated  StatusDeal = "INITIATED"
	InProgress StatusDeal = "IN-PROGRESS"
	ClosedWon  StatusDeal = "CLOSED-WON"
	ClosedLost StatusDeal = "CLOSED-LOST"
)

func (sd *StatusDeal) Scan(value interface{}) error {
	*sd = StatusDeal(value.([]byte))
	return nil
}

func (sd StatusDeal) Value() (driver.Value, error) {
	return string(sd), nil
}

type Deal struct {
	gorm.Model
	Title string `gorm:"varchar(255);not null" json:"title"`
	Value uint   `gorm:"varchar(255);not null" json:"value"`
	//Status StatusDeal `gorm:"column:status_deal; type:ENUM('INITIATED', 'IN-PROGRESS', 'CLOSED-WON', 'CLOSED-LOST')" json:"status"`
	Status    StatusDeal `gorm:"type:status_deal"`
	ContactID uint       `json:"contact_id"`
	RepID     uint       `json:"rep_id"`
}
