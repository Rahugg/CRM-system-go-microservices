package service

import (
	"crm_system/internal/auth/controller/consumer/dto"
	entity2 "crm_system/internal/auth/entity"
	"crm_system/pkg/auth/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
)

func (s *Service) SignUp(ctx *gin.Context, payload *entity2.SignUpInput, roleId uint, provider string) (string, error) {
	if !utils.IsValidEmail(payload.Email) {
		return "", errors.New("email validation error")
	}

	if !utils.IsValidPassword(payload.Password) {
		return "", errors.New("passwords should contain:\nUppercase letters: A-Z\nLowercase letters: a-z\nNumbers: 0-9")
	}

	if payload.Password != payload.PasswordConfirm {
		return "", errors.New("passwords do not match")
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return "", err
	}

	user := entity2.User{
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		Email:       strings.ToLower(payload.Email),
		Password:    hashedPassword,
		RoleID:      roleId,
		Provider:    provider,
		IsConfirmed: false,
	}

	if err = s.Repo.CreateUser(ctx, &user); err != nil {
		return "", err
	}
	userResponse, err := s.Repo.GetUserByEmail(ctx, user.Email)

	return userResponse.ID.String(), nil
}

func (s *Service) SignIn(ctx *gin.Context, payload *entity2.SignInInput) (*entity2.SignInResult, error) {
	user, err := s.Repo.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	if err = utils.VerifyPassword(user.Password, payload.Password); err != nil {
		return nil, err
	}

	accessToken, err := utils.GenerateToken(s.Config.Jwt.AccessTokenExpiredIn, user.ID, s.Config.Jwt.AccessPrivateKey)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateToken(s.Config.Jwt.RefreshTokenExpiredIn, user.ID, s.Config.Jwt.RefreshPrivateKey)
	if err != nil {
		return nil, err
	}

	data := entity2.SignInResult{
		Role:            user.Role.Name,
		AccessToken:     accessToken,
		RefreshToken:    refreshToken,
		AccessTokenAge:  int(s.Config.Jwt.AccessTokenMaxAge * 60),
		RefreshTokenAge: int(s.Config.Jwt.RefreshTokenMaxAge * 60),
	}
	return &data, nil
}

func (s *Service) returnUsers(ctx *gin.Context, roleId uint) (*[]entity2.SignUpResult, error) {
	users, err := s.Repo.GetUsersByRole(ctx, roleId)
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

func (s *Service) GetUsers(ctx *gin.Context, sortBy, sortOrder, age string) (*[]entity2.User, error) {
	users, err := s.Repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	if age != "" {
		users, err = s.filterUsersByAge(users, age)
	}

	if sortBy != "" {
		users, err = s.sortUsers(users, sortBy, sortOrder)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func (s *Service) sortUsers(users *[]entity2.User, sortBy, sortOrder string) (*[]entity2.User, error) {
	users, err := s.Repo.SortUsers(users, sortBy, sortOrder)

	if err != nil {
		return nil, err
	}

	return users, nil
}
func (s *Service) filterUsersByAge(users *[]entity2.User, age string) (*[]entity2.User, error) {
	users, err := s.Repo.FilterUsersByAge(users, age)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) GetUser(ctx *gin.Context, id string) (*entity2.User, error) {
	user, err := s.Repo.GetUser(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *Service) CreateUser(ctx *gin.Context, user *entity2.User) error {
	if err := s.Repo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
func (s *Service) UpdateUser(ctx *gin.Context, newUser *entity2.User, id string) error {
	user, err := s.Repo.GetUser(id)
	if err != nil {
		return err
	}

	if newUser.FirstName != "" {
		user.FirstName = newUser.FirstName
	}

	if newUser.LastName != "" {
		user.LastName = newUser.LastName
	}

	if newUser.Age != 0 {
		user.Age = newUser.Age
	}

	if newUser.Phone != "" {
		user.Phone = newUser.Phone
	}

	if newUser.RoleID != 0 {
		user.RoleID = newUser.RoleID
	}

	if newUser.Email != "" {
		user.Email = newUser.Email
	}

	if newUser.Provider != "" {
		user.Provider = newUser.Provider
	}

	if newUser.Password != "" {
		user.Password = newUser.Password
	}

	if err = s.Repo.SaveUser(user); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUser(ctx *gin.Context, id string) error {
	user, err := s.Repo.GetUser(id)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteUser(ctx, id, user); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetMe(ctx *gin.Context) (interface{}, error) {
	user, exist := ctx.MustGet("currentUser").(*entity2.User)
	if !exist {
		return nil, fmt.Errorf("the current user did not authorize or does not exist")
	}

	return user, nil
}

func (s *Service) UpdateMe(ctx *gin.Context, newUser *entity2.User) error {
	user := ctx.MustGet("currentUser").(*entity2.User)

	if newUser.FirstName != "" {
		user.FirstName = newUser.FirstName
	}

	if newUser.LastName != "" {
		user.LastName = newUser.LastName
	}

	if newUser.Age != 0 {
		user.Age = newUser.Age
	}

	if newUser.Phone != "" {
		user.Phone = newUser.Phone
	}

	if newUser.RoleID != 0 {
		user.RoleID = newUser.RoleID
	}

	if newUser.Email != "" {
		user.Email = newUser.Email
	}

	if newUser.Provider != "" {
		user.Provider = newUser.Provider
	}

	if newUser.Password != "" {
		user.Password = newUser.Password
	}

	if err := s.Repo.SaveUser(user); err != nil {
		return err
	}

	return nil
}

func (s *Service) SearchUser(ctx *gin.Context, query string) (*[]entity2.User, error) {
	users, err := s.Repo.SearchUser(ctx, query)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *Service) CreateUserCode(ctx *gin.Context, id string) error {
	randNum1 := rand.Intn(999-100) + 100
	randNum2 := rand.Intn(999-100) + 100

	userCode := dto.UserCode{Code: fmt.Sprintf("%d%d", randNum1, randNum2)}
	b, err := json.Marshal(&userCode)
	if err != nil {
		return fmt.Errorf("failed to marshall UserCode err: %w", err)
	}

	s.userVerificationProducer.ProduceMessage(b)
	err = s.Repo.CreateUserCode(id, userCode.Code)
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) ConfirmUser(ctx *gin.Context, code string) error {
	err := s.Repo.ConfirmUser(code)
	if err != nil {
		return err
	}

	return nil

}
