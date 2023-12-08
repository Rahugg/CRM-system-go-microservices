package grpc

import (
	"context"
	"crm_system/config/auth"
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
	for _, Role := range request.Roles {
		if role.Name == Role || Role == "any" {
			response := &pb.ValidateResponse{
				Response: &pb.ResponseJSON{
					User: &pb.User{
						Id:        cast.ToString(user.ID),
						FirstName: user.FirstName,
						LastName:  user.LastName,
						Age:       int64(user.Age),
						Phone:     user.Phone,
						RoleID:    int64(user.RoleID),
						Email:     user.Email,
						Provider:  user.Provider,
						Password:  user.Password,
					},
					Role: &pb.Role{
						ID:   int64(role.ID),
						Name: role.Name,
					},
				},
			}
			return response, nil
		}
	}
	return nil, err
}
