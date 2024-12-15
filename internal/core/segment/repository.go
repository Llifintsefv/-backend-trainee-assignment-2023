package segment

import (
	"backend-trainee-assignment-2023/internal/core/interfaces"
	"context"
	"database/sql"

	"golang.org/x/net/context/ctxhttp"
)


type segmentRepository struct {
	db *sql.DB
	
}

func NewSegmentRepository(db *sql.DB)  interfaces.SegmentRepository {
	return &segmentRepository{db: db}
}


func (r *segmentRepository) CreateSegment (ctx context.Context,slug string,AutoAddPercent int) (int,error) {
	var id int
	err := r.db.QueryRowContext(ctx,"INSERT INTO segments (slug,auto_add_percent) VALUES ($1,$2) RETURNING id",slug,AutoAddPercent).Scan(&id)
	if err != nil {
		return 0,err
	}

	return id,nil
}

func (r *segmentRepository) DeleteSegment(ctx context.Context,slug string) error {
	_,err := r.db.ExecContext(ctx,"DELETE FROM segments WHERE slug = $1",slug)
	if err != nil {
		return err
	}
	return nil
}


func (r *segmentRepository) CreateUserSegment(ctx context.Context,userId int,segmentId int) error {
	_,err := r.db.ExecContext(ctx,"INSERT INTO user_segments (user_id,segment_id) VALUES ($1,$2)",userId,segmentId)
	if err != nil {
		return err
	}
	return nil
}




