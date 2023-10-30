package service

import (
	entity2 "crm_system/internal/auth/entity"
	"crm_system/pkg/auth/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func (s *Service) SignUp(ctx *gin.Context, payload *entity2.SignUpInput, roleId uint, verified bool, provider string) (interface{}, error) {
	if !utils.IsValidEmail(payload.Email) {
		return nil, errors.New("email validation error")
	}

	if !utils.IsValidPassword(payload.Password) {
		return nil, errors.New("passwords should contain:\nUppercase letters: A-Z\nLowercase letters: a-z\nNumbers: 0-9")
	}

	if payload.Password != payload.PasswordConfirm {
		return nil, errors.New("passwords do not match")
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	user := entity2.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		RoleID:    roleId,
		Provider:  provider,
	}

	if err = s.Repo.CreateUser(&user); err != nil {
		return nil, err
	}

	result, err := s.returnUsers(user.RoleID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Service) SignIn(payload *entity2.SignInInput) (*entity2.SignInResult, error) {
	user, err := s.Repo.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err = utils.VerifyPassword(user.Password, payload.Password); err != nil {
		return nil, err
	}

	accessToken, err := utils.GenerateToken(time.Duration(s.Config.Jwt.AccessTokenExpiredIn), user.ID, s.Config.Jwt.AccessPrivateKey)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateToken(time.Duration(s.Config.Jwt.RefreshTokenExpiredIn), user.ID, s.Config.Jwt.RefreshPrivateKey)
	if err != nil {
		return nil, err
	}

	data := entity2.SignInResult{
		Role:         user.Role.Name,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		//AccessTokenAge:  int(s.Config.Jwt.AccessTokenMaxAge * 60),
		//RefreshTokenAge: int(s.Config.Jwt.RefreshTokenMaxAge * 60),
	}
	return &data, nil
}

func (s *Service) returnUsers(roleId uint) (*[]entity2.SignUpResult, error) {
	users, err := s.Repo.GetUsersByRole(roleId)
	if err != nil {
		return nil, err
	}

	result := make([]entity2.SignUpResult, len(*users))

	for i, user := range *users {
		result[i] = entity2.SignUpResult{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role.Name,
			Provider:  user.Provider,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}
	return &result, nil
}
