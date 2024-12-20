package segment

import (
	"backend-trainee-assignment-2023/internal/core/interfaces"
	"backend-trainee-assignment-2023/internal/core/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
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
	tx,err := r.db.BeginTx(ctx,nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	 _, err = tx.ExecContext(ctx, "UPDATE user_segments SET deleted_at = CURRENT_TIMESTAMP WHERE slug = $1 AND deleted_at IS NULL", slug)
    if err != nil {
        return fmt.Errorf("failed to update user_segments: %w", err)
    }

	_,err = tx.ExecContext(ctx,"DELETE FROM segments WHERE slug = $1",slug)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}

func(r *segmentRepository) SegmentExists(id int) (bool,error) {
	var exists bool
	err := r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM segments WHERE id = $1)`).Scan(&exists)
	if err != nil {
		return false,err
	}
	return true,nil
}

func (r *segmentRepository) GetSegmentIdBySlug(ctx context.Context,slug string) (int,error) {
	var id int
	err := r.db.QueryRowContext(ctx,"SELECT id FROM segments WHERE slug = $1",slug).Scan(&id)
	if err != nil {
		return 0,err
	}
	return id,nil
}

func (r *segmentRepository) DeleteUserSegment(ctx context.Context, userId int, segmentId int) error {
	_, err := r.db.ExecContext(ctx, "UPDATE user_segments SET deleted_at = $1 WHERE user_id = $2 AND segment_id = $3 AND deleted_at IS NULL", time.Now(), userId, segmentId)
	return err
}

func (r *segmentRepository) CreateUserSegment(ctx context.Context,userId int,segmentId int,ExpiresAt time.Time) error {
	tx,err := r.db.BeginTx(ctx,nil)
	if err != nil {

	}
	log.Println(userId,segmentId,ExpiresAt)
	defer tx.Rollback()

	var exists bool

	err = tx.QueryRowContext(ctx,"SELECT EXISTS(SELECT 1 FROM user_segments WHERE user_id = $1 AND segment_id = $2)").Scan(exists)
	if err != nil {
		return err
	}
	if exists {
		if !ExpiresAt.IsZero() {
			_,err := tx.ExecContext(ctx,"UPDATE user_segments SET expires_at = $1 WHERE user_id = $2 AND segment_id = $3",userId,segmentId,ExpiresAt)
			if err != nil {

			}
			return tx.Commit()
		}
	}

	_,err = tx.ExecContext(ctx,"INSERT INTO user_segments (user_id,segment_id,deleted_at) VALUES ($1,$2,$3)",userId,segmentId,ExpiresAt)
	if err != nil {
		return err
	}
	log.Println(userId,segmentId,ExpiresAt)
	return tx.Commit()
}

func (r *segmentRepository) GetUserSegments(ctx context.Context,userId int) ([]models.GetUserSegmentsResponse,error) {
	rows,err := r.db.QueryContext(ctx, `
		SELECT us.segment_id, us.created_at, us.deleted_at, s.slug
        FROM user_segments us
        JOIN segments s ON us.segment_id = s.id
        WHERE us.user_id = $1 AND us.deleted_at IS NULL`,userId)
	if err != nil {
		return nil,err
	}

	defer rows.Close()

	var segments []models.GetUserSegmentsResponse
	for rows.Next(){
		var us models.GetUserSegmentsResponse
		var deletedAt sql.NullTime
		if err := rows.Scan(&us.SegmentId,&us.CreatedAt,&deletedAt,&us.Slug); err != nil {
			return nil,err
		}
		if deletedAt.Valid {
			us.DeletedAt = &deletedAt.Time
		} else {
			us.DeletedAt = nil
		}
		segments = append(segments, us)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return segments, nil
}




