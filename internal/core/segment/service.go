package segment

import (
	"backend-trainee-assignment-2023/internal/core/interfaces"
	"backend-trainee-assignment-2023/internal/core/models"
	"context"
	"fmt"
	"time"
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

func (s *segmentService) CreateUserSegment(ctx context.Context,userId int,Add []string, Remove []string, TTL map[string]string) (error) {
	_,err := s.userRepo.UserExists(userId)
	if err != nil {
		
	}

	for _,slug := range Add {
		segmentId,err := s.segmentRepo.GetSegmentIdBySlug(ctx,slug)
		if err != nil {
			return fmt.Errorf("segment %s not found",slug)
		}
		var ExpiresAt time.Time
		if ttlStr,ok := TTL[slug]; ok {
			ExpiresAt,err = time.Parse(time.RFC3339,ttlStr)
			if err != nil {
				return err
			}
		}

		err = s.segmentRepo.CreateUserSegment(ctx,userId,segmentId,ExpiresAt)
		if err != nil {
			return err
		}
	}

	for _,slug := range Remove {
		segmentId,err := s.segmentRepo.GetSegmentIdBySlug(ctx,slug)
		if err != nil {
			return fmt.Errorf("segment %s not found",slug)
		}
		}
	}
	
	
	// segmentId,err := s.segmentRepo.GetSegmentIdBySlug(ctx,userSegment.Slug)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return err
	// 	} else {
	// 		return err
	// 	}
	// }

	
	// ExpiresAt,err := time.Parse(time.RFC3339,userSegment.ExpiresAt)
	// if err != nil {

	// }

	// err = s.segmentRepo.CreateUserSegment(ctx,userSegment.UserId,segmentId,ExpiresAt)
	// if err != nil {

	// }

	// return nil



}