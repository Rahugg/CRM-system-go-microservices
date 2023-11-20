package entity

import (
	"github.com/google/uuid"
	"time"
)

type SignUpInput struct {
	FirstName       string `gorm:"type:varchar(255);not null" json:"name" binding:"required"`
	LastName        string `gorm:"type:varchar(255);not null" json:"surname" binding:"required"`
	Email           string `gorm:"type:varchar(255);not null" json:"email" binding:"required"`
	Password        string `gorm:"type:varchar(255);not null" json:"password" binding:"required,min=8"`
	PasswordConfirm string `gorm:"type:varchar(255);not null" json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Email    string `gorm:"type:varchar(255);not null" json:"email"  binding:"required"`
	Password string `gorm:"type:varchar(255);not null" json:"password"  binding:"required"`
}

type SignInResult struct {
	Role            string `json:"role"`
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	AccessTokenAge  int    `json:"access_token_age"`
	RefreshTokenAge int    `json:"refresh_token_age"`
}

type SignUpResult struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"name"`
	LastName  string    `json:"surname"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CustomResponseWithData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
