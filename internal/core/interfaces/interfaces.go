package interfaces

import (
	"backend-trainee-assignment-2023/internal/core/models"
	"context"
	"time"
)

type SegmentRepository interface {
	CreateSegment(ctx context.Context, slug string, AutoAddPercent int) (int, error)
	SegmentExists(id int) (bool,error)
	GetSegmentIdBySlug(ctx context.Context, slug string) (int, error)
	DeleteSegment(ctx context.Context, slug string) (error)
	DeleteUserSegment(ctx context.Context, userId int, segmentId int) error
	CreateUserSegment(ctx context.Context, userId int, segmentId int, ExpiresAt time.Time) error
}

type SegmentService interface {
	CreateSegment(ctx context.Context, segment models.Segment) (int, error)
	CreateUserSegment(ctx context.Context,userId int, Add []string, Remove []string, TTL map[string]string) (error)
	DeleteSegment(ctx context.Context, slug string) (error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, id int) error
	UserExists(id int) (bool,error)
}

type UserService interface {
	CreateUser(ctx context.Context, user models.User) error
}
