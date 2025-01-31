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

func (s *segmentService) DeleteSegment(ctx context.Context, slug string) (error) {
	segmentId,err := s.segmentRepo.GetSegmentIdBySlug(ctx,slug)
	if err != nil {
		return fmt.Errorf("segment %s not found",slug)
	}
	_,err = s.segmentRepo.SegmentExists(segmentId)
	
	if err != nil {
		return fmt.Errorf("segment %s not found",slug)
	}

	err = s.segmentRepo.DeleteSegment(ctx,slug)
	if err != nil {
		return fmt.Errorf("failed to delete segment %s",slug)	
	}
	return nil
}

func (s *segmentService) DeleteUserSegment(ctx context.Context,userId int,segmentId int) (error) {
	return s.segmentRepo.DeleteUserSegment(ctx,userId,segmentId)
}

func (s *segmentService) CreateUserSegment(ctx context.Context,userId int,Add []string, Remove []string, TTL map[string]string) (error) {
	_,err := s.userRepo.UserExists(userId)
	if err != nil {
		return fmt.Errorf("user %d not found",userId)
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
		err = s.segmentRepo.DeleteUserSegment(ctx,userId,segmentId)
		if err != nil { 
			return fmt.Errorf("failed to delete segment %s",slug)
		}
		}
		return nil
	}
	
	
func (s *segmentService) GetUserSegments(ctx context.Context,userId int) ([]models.GetUserSegmentsResponse, error) {
	exists, err := s.userRepo.UserExists(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("user %d not found", userId)
	}
	
	UserSegments,err := s.segmentRepo.GetUserSegments(ctx,userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user segments: %w", err)
	}

	return UserSegments,nil

}