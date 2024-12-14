package user

import (
	"backend-trainee-assignment-2023/internal/core/interfaces"
	"backend-trainee-assignment-2023/internal/core/models"
	"context"
)



type userService struct {
	userRepo    interfaces.UserRepository
	segmentRepo	interfaces.SegmentRepository
}

func NewUserService(userRepo  interfaces.UserRepository, segmentRepo  interfaces.SegmentRepository)  interfaces.UserService {
	return &userService{userRepo: userRepo, segmentRepo: segmentRepo}
}

func (s *userService) CreateUser(ctx context.Context,user models.User) error {
	return s.userRepo.CreateUser(ctx,user.Id)
}