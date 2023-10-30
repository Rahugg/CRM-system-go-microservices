package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Address   string    `gorm:"type:varchar(255);not null" json:"address"`
	Phone     string    `gorm:"type:varchar(255);not null" json:"phone"`
	ManagerID uuid.UUID `json:"manager_id"`
}

type NewCompany struct {
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
	Address string `gorm:"type:varchar(255);not null" json:"address"`
	Phone   string `gorm:"type:varchar(255);not null" json:"phone"`
}
