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

func getUser(payload *entity2.SignUpInput, hashedPassword, provider string, roleId uint) entity2.User {
	userBuilder := entity2.NewUser()
	user := userBuilder.
		SetFirstName(payload.FirstName).
		SetLastName(payload.LastName).
		SetEmail(strings.ToLower(payload.Email)).
		SetPassword(hashedPassword).
		SetRoleID(roleId).
		SetProvider(provider).
		Build()
	return user
}

func (s *Service) SignUp(payload *entity2.SignUpInput, roleId uint, provider string) (string, error) {
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

	user := getUser(payload, hashedPassword, provider, roleId)

	if err = s.Repo.CreateUser(&user); err != nil {
		return "", err
	}

	userResponse, err := s.Repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	return userResponse.ID.String(), nil
}

func (s *Service) SignIn(payload *entity2.SignInInput) (*entity2.SignInResult, error) {
	user, err := s.Repo.GetUserByEmail(payload.Email)
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

func (s *Service) GetUsers(sortBy, sortOrder, age string) (*[]entity2.User, error) {
	users, err := s.Storage.GetAllUsers()
	if err != nil {
		return nil, err
	}
	if age != "" {
		users, err = s.filterUsersByAge(users, age)
		if err != nil {
			return nil, err
		}
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

func (s *Service) GetUser(id string) (*entity2.User, error) {
	user, err := s.Repo.GetUser(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *Service) CreateUser(user *entity2.User) error {
	if err := s.Repo.CreateUser(user); err != nil {
		return err
	}

	return nil
}
func (s *Service) UpdateUser(newUser *entity2.User, id string) error {
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

func (s *Service) DeleteUser(id string) error {
	user, err := s.Repo.GetUser(id)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteUser(id, user); err != nil {
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

func (s *Service) SearchUser(query string) (*[]entity2.User, error) {
	users, err := s.Repo.SearchUser(query)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *Service) CreateUserCode(id string) error {
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

func (s *Service) ConfirmUser(code string) error {
	err := s.Repo.ConfirmUser(code)
	if err != nil {
		return err
	}

	return nil

}
