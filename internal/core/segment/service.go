package segment

import (
	"backend-trainee-assignment-2023/internal/core/interfaces"
	"backend-trainee-assignment-2023/internal/core/models"
	"context"
)


type segmentService struct {
	segmentRepo  interfaces.SegmentRepository
	userRepo     interfaces.UserRepository
}

func NewSegmentService(segmentRepo  interfaces.SegmentRepository, userRepo  interfaces.UserRepository)  interfaces.SegmentService {
	return &segmentService{segmentRepo: segmentRepo, userRepo: userRepo}
}


func (s *segmentService) CreateSegment(ctx context.Context,segment models.Segment) (int,error) {
	return s.segmentRepo.CreateSegment(ctx,segment.Slug,segment.AutoAddPercent)
}