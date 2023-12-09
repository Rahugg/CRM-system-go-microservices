package grpc

import (
	"context"
	"crm_system/config/auth"
	"crm_system/internal/auth/entity"
	"crm_system/internal/auth/repository"
	pb "crm_system/pkg/auth/authservice/gw"
	"crm_system/pkg/auth/logger"
	"crm_system/pkg/auth/utils"
	"fmt"
	"github.com/spf13/cast"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	logger *logger.Logger
	repo   *repository.AuthRepo
	config *auth.Configuration
}

func NewService(logger *logger.Logger, repo *repository.AuthRepo, config *auth.Configuration) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
		config: config,
	}
}

func getUser(user *entity.User) *pb.User {
	return &pb.User{
		Id:        cast.ToString(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       int64(user.Age),
		Phone:     user.Phone,
		RoleID:    int64(user.RoleID),
		Email:     user.Email,
		Provider:  user.Provider,
		Password:  user.Password,
	}
}
func getRole(role *entity.Role) *pb.Role {
	return &pb.Role{
		ID:   int64(role.ID),
		Name: role.Name,
	}
}

func getResponse(request *pb.ValidateRequest, role *entity.Role, user *entity.User) (*pb.ValidateResponse, error) {
	for _, Role := range request.Roles {
		if role.Name == Role || Role == "any" {
			response := &pb.ValidateResponse{
				Response: &pb.ResponseJSON{
					User: getUser(user),
					Role: getRole(role),
				},
			}
			return response, nil
		}
	}
	return nil, fmt.Errorf("error happened in getting a Validate Response")
}

func (s *Service) Validate(ctx context.Context, request *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	sub, err := utils.ValidateToken(request.AccessToken, s.config.Jwt.AccessPrivateKey)

	if err != nil {
		s.logger.Error("failed to ValidateToken err %v", err)
		return nil, err
	}

	user, err := s.repo.GetUserByIdWithPreload(fmt.Sprint(sub))
	if err != nil {
		return nil, fmt.Errorf("the user belonging to this token no logger exists")
	}

	role, _ := s.repo.GetRoleById(user.RoleID)

	return getResponse(request, role, user)
}
