package interfaces

import (
	"backend-trainee-assignment-2023/internal/core/models"
	"context"
)

type SegmentRepository interface {
	CreateSegment(ctx context.Context, slug string, AutoAddPercent int) (int, error)
}

type SegmentService interface {
	CreateSegment(ctx context.Context, segment models.Segment) (int, error)
}

type UserRepository interface {
}

type UserService interface {
}
