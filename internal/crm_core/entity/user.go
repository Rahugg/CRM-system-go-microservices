package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	FirstName string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(255);not null" json:"last_name"`
	Age       uint64    `json:"age"`
	Phone     string    `gorm:"type:varchar(255);" json:"phone"`
	RoleID    uint      `gorm:"not null" json:"role_id"`
	Role      Role      `gorm:"not null" json:"role"`
	Email     string    `gorm:"type:varchar(255);not null;uniqueIndex;not null" json:"email"`
	Provider  string    `gorm:"type:varchar(255);not null" json:"provider"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
}
