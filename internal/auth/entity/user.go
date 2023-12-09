package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"ID"`
	FirstName   string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName    string    `gorm:"type:varchar(255);not null" json:"last_name"`
	Age         uint64    `json:"age"`
	Phone       string    `gorm:"type:varchar(255);" json:"phone"`
	RoleID      uint      `gorm:"not null" json:"role_id"`
	Role        Role      `gorm:"not null" json:"role"`
	Email       string    `gorm:"type:varchar(255);not null;uniqueIndex;not null" json:"email"`
	Provider    string    `gorm:"type:varchar(255);not null" json:"provider"`
	Password    string    `gorm:"type:varchar(255);not null" json:"-"`
	IsConfirmed bool      `json:"is_confirmed"`
}

type UserCode struct {
	UserID string `json:"user_id"`
	Code   string `json:"code"`
}
type InputCode struct {
	Code string `json:"code"`
}

type UserBuilderI interface {
	SetGORMModel(val gorm.Model) UserBuilderI
	SetID(val uuid.UUID) UserBuilderI
	SetFirstName(val string) UserBuilderI
	SetLastName(val string) UserBuilderI
	SetAge(val uint64) UserBuilderI
	SetPhone(val string) UserBuilderI
	SetRoleID(val uint) UserBuilderI
	SetRole(val Role) UserBuilderI
	SetEmail(val string) UserBuilderI
	SetProvider(val string) UserBuilderI
	SetPassword(val string) UserBuilderI
	SetIsConfirmed(val bool) UserBuilderI
	Build() User
}

func NewUser() UserBuilderI {
	return User{}.SetGORMModel(gorm.Model{}).SetID(uuid.UUID{}).SetFirstName("firstName").
		SetLastName("lastName").
		SetEmail("email@gmail.com").
		SetAge(18).
		SetPhone("87777771234").
		SetRoleID(2).
		SetRole(Role{}).
		SetProvider("manager").
		SetPassword("123456789aA").
		SetIsConfirmed(false).
		Build()
}

func (u User) SetGORMModel(val gorm.Model) UserBuilderI {
	u.Model = val
	return u
}
func (u User) SetID(val uuid.UUID) UserBuilderI {
	u.ID = val
	return u
}

func (u User) SetFirstName(val string) UserBuilderI {
	u.FirstName = val
	return u
}
func (u User) SetLastName(val string) UserBuilderI {
	u.LastName = val
	return u
}
func (u User) SetAge(val uint64) UserBuilderI {
	u.Age = val
	return u
}
func (u User) SetPhone(val string) UserBuilderI {
	u.Phone = val
	return u
}
func (u User) SetRoleID(val uint) UserBuilderI {
	u.RoleID = val
	return u
}
func (u User) SetRole(val Role) UserBuilderI {
	u.Role = val
	return u
}
func (u User) SetEmail(val string) UserBuilderI {
	u.Email = val
	return u
}
func (u User) SetProvider(val string) UserBuilderI {
	u.Provider = val

	return u
}
func (u User) SetPassword(val string) UserBuilderI {
	u.Password = val
	return u
}

func (u User) SetIsConfirmed(val bool) UserBuilderI {
	u.IsConfirmed = val
	return u
}

func (u User) Build() User {
	return User{
		Model:       u.Model,
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Age:         u.Age,
		Phone:       u.Phone,
		RoleID:      u.RoleID,
		Role:        u.Role,
		Email:       u.Email,
		Provider:    u.Provider,
		Password:    u.Password,
		IsConfirmed: u.IsConfirmed,
	}
}
