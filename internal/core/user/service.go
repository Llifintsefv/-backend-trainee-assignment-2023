package user

import (
	"backend-trainee-assignment-2023/internal/core/interfaces"
)



type userService struct {
	userRepo    interfaces.UserRepository
	segmentRepo	interfaces.SegmentRepository
}

func NewUserService(userRepo  interfaces.UserRepository, segmentRepo  interfaces.SegmentRepository)  interfaces.UserService {
	return &userService{userRepo: userRepo, segmentRepo: segmentRepo}
}

