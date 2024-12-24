package segment

import (
	"backend-trainee-assignment-2023/internal/core/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func TestMain(m *testing.M) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbSSLMode == "" {
		log.Fatal("DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL_MODE must be set")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Run tests
	code := m.Run()

	os.Exit(code)
}

func setupDB(t *testing.T) {
	// Очистка данных перед каждым тестом
	_, err := db.Exec("DELETE FROM user_segments")
	if err != nil {
		t.Fatalf("Failed to clear user_segments table: %v", err)
	}
	_, err = db.Exec("DELETE FROM segments")
	if err != nil {
		t.Fatalf("Failed to clear segments table: %v", err)
	}
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatalf("Failed to clear users table: %v", err)
	}
}

func TestSegmentRepository_CreateSegment(t *testing.T) {
	setupDB(t)
	repo := NewSegmentRepository(db)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		id, err := repo.CreateSegment(ctx, "AVITO_VOICE_MESSAGES", 0)
		if err != nil {
			t.Fatalf("Failed to create segment: %v", err)
		}

		if id <= 0 {
			t.Errorf("Expected id > 0, got %d", id)
		}
		var slug string
		var AutoAddPercent int
		err = db.QueryRow("SELECT slug, auto_add_percent FROM segments WHERE id = $1", id).Scan(&slug,&AutoAddPercent)
		if err != nil {
			t.Fatalf("Failed to get segment from DB: %v", err)
		}

		if slug != "AVITO_VOICE_MESSAGES" {
			t.Errorf("Expected slug 'AVITO_VOICE_MESSAGES', got '%s'", slug)
		}
		if AutoAddPercent != 0 {
			t.Errorf("Expected percent '0' , got '%d'", AutoAddPercent)
		}
	})

	t.Run("Duplicate slug", func(t *testing.T) {
		_, err := repo.CreateSegment(ctx, "AVITO_VOICE_MESSAGES", 0)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestSegmentRepository_DeleteSegment(t *testing.T) {
	setupDB(t)
	repo := NewSegmentRepository(db)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		id, err := repo.CreateSegment(ctx, "AVITO_VOICE_MESSAGES", 0)
		if err != nil {
			t.Fatalf("Failed to create segment: %v", err)
		}

		err = repo.DeleteSegment(ctx, "AVITO_VOICE_MESSAGES")
		if err != nil {
			t.Fatalf("Failed to delete segment: %v", err)
		}

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM segments WHERE id = $1", id).Scan(&count)
		if err != nil {
			t.Fatalf("Failed to query segment: %v", err)
		}

		if count != 0 {
			t.Errorf("Expected segment to be deleted, but it still exists")
		}
	})

	t.Run("Not found", func(t *testing.T) {
		err := repo.DeleteSegment(ctx, "NON_EXISTENT_SEGMENT")
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})
}

func TestSegmentRepository_SegmentExists(t *testing.T) {
	setupDB(t)
	repo := NewSegmentRepository(db)
	ctx := context.Background()

	segmentId, err := repo.CreateSegment(ctx, "AVITO_VOICE_MESSAGES", 0)
	if err != nil {
		t.Fatalf("Failed to create segment: %v", err)
	}

	t.Run("Exists", func(t *testing.T) {
		exists, err := repo.SegmentExists(segmentId)
		if err != nil {
			t.Fatalf("Failed to check if segment exists: %v", err)
		}

		if !exists {
			t.Errorf("Expected segment to exist, but it doesn't")
		}
	})

	t.Run("Not exists", func(t *testing.T) {
		exists, err := repo.SegmentExists(9999)
		if err != nil {
			t.Fatalf("Failed to check if segment exists: %v", err)
		}

		if exists {
			t.Errorf("Expected segment to not exist, but it does")
		}
	})
}

func TestSegmentRepository_GetSegmentIdBySlug(t *testing.T) {
	setupDB(t)
	repo := NewSegmentRepository(db)
	ctx := context.Background()

	expectedId, err := repo.CreateSegment(ctx, "AVITO_VOICE_MESSAGES", 0)
	if err != nil {
		t.Fatalf("Failed to create segment: %v", err)
	}

	t.Run("Success", func(t *testing.T) {
		id, err := repo.GetSegmentIdBySlug(ctx, "AVITO_VOICE_MESSAGES")
		if err != nil {
			t.Fatalf("Failed to get segment ID by slug: %v", err)
		}

		if id != expectedId {
			t.Errorf("Expected ID %d, got %d", expectedId, id)
		}
	})

	t.Run("Not found", func(t *testing.T) {
		_, err := repo.GetSegmentIdBySlug(ctx, "NON_EXISTENT_SEGMENT")
		if err != sql.ErrNoRows {
			t.Errorf("Expected sql.ErrNoRows, got: %v", err)
		}
	})
}

func TestSegmentRepository_DeleteUserSegment(t *testing.T) {
	setupDB(t)
	repo := NewSegmentRepository(db)
	ctx := context.Background()
	_,err := db.ExecContext(ctx,"INSERT INTO users (id) VALUES (1)")
	if err != nil {
		t.Fatal(err)
	}
	segmentId,err := repo.CreateSegment(ctx,"AVITO_VOICE_MESSAGES", 0)
	if err != nil {
		t.Fatal(err)
	}
	_,err = db.ExecContext(ctx,"INSERT INTO user_segments (user_id,segment_id) VALUES (1,$1)",segmentId)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Success", func(t *testing.T) {
		err := repo.DeleteUserSegment(ctx, 1, segmentId)
		if err != nil {
			t.Fatalf("Failed to delete user segment: %v", err)
		}

		var deletedAt sql.NullTime
		err = db.QueryRow("SELECT deleted_at FROM user_segments WHERE user_id = 1 AND segment_id = $1", segmentId).Scan(&deletedAt)
		if err != nil {
			t.Fatalf("Failed to query user segment: %v", err)
		}

		if !deletedAt.Valid {
			t.Errorf("Expected deleted_at to be set, but it's not")
		}
	})

	t.Run("User segment not found", func(t *testing.T) {
		err := repo.DeleteUserSegment(ctx, 1, 9999) 
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})
}

func TestSegmentRepository_CreateUserSegment(t *testing.T) {
	setupDB(t)
	repo := NewSegmentRepository(db)
	ctx := context.Background()

	userId := 1
	segmentId := 1
	_,err := db.ExecContext(ctx,"INSERT INTO users (id) VALUES ($1)",userId)
	if err != nil {
		t.Fatal(err)
	}
	_,err = repo.CreateSegment(ctx,"AVITO_VOICE_MESSAGES", 0)
	if err != nil {
		t.Fatal(err)
	}
	

	t.Run("Success", func(t *testing.T) {
		expiresAt := time.Now().Add(time.Hour)
		err := repo.CreateUserSegment(ctx, userId, segmentId, expiresAt)
		if err != nil {
			t.Fatalf("Failed to create user segment: %v", err)
		}

		var createdAt, deletedAt sql.NullTime
		var dbExpiresAt time.Time
		err = db.QueryRow("SELECT created_at, deleted_at, expires_at FROM user_segments WHERE user_id = $1 AND segment_id = $2", userId, segmentId).Scan(&createdAt, &deletedAt, &dbExpiresAt)
		if err != nil {
			t.Fatalf("Failed to query user segment: %v", err)
		}

		if !createdAt.Valid {
			t.Errorf("Expected created_at to be set, but it's not")
		}
		if deletedAt.Valid {
			t.Errorf("Expected deleted_at to be null, but it's set")
		}
		if !dbExpiresAt.Equal(expiresAt) {
			t.Errorf("Expected expires_at to be %v, got %v", expiresAt, dbExpiresAt)
		}
	})

	t.Run("Success without TTL", func(t *testing.T) {
		err := repo.CreateUserSegment(ctx, userId, segmentId, time.Time{})
		if err != nil {
			t.Fatalf("Failed to create user segment: %v", err)
		}
		var createdAt, deletedAt sql.NullTime
		var dbExpiresAt sql.NullTime
		err = db.QueryRow("SELECT created_at, deleted_at, expires_at FROM user_segments WHERE user_id = $1 AND segment_id = $2 AND expires_at IS NULL", userId, segmentId).Scan(&createdAt, &deletedAt, &dbExpiresAt)
		if err != nil {
			t.Fatalf("Failed to query user segment: %v", err)
		}
	
		if !createdAt.Valid {
			t.Errorf("Expected created_at to be set, but it's not")
		}
		if deletedAt.Valid {
			t.Errorf("Expected deleted_at to be null, but it's set")
		}
		if dbExpiresAt.Valid {
			t.Errorf("Expected expires_at to be null, but it's set")
		}
	})

	t.Run("Update TTL", func(t *testing.T) {
		newExpiresAt := time.Now().Add(2 * time.Hour)
		err := repo.CreateUserSegment(ctx, userId, segmentId, newExpiresAt)
		if err != nil {
			t.Fatalf("Failed to update user segment: %v", err)
		}

		var dbExpiresAt time.Time
		err = db.QueryRow("SELECT expires_at FROM user_segments WHERE user_id = $1 AND segment_id = $2", userId, segmentId).Scan(&dbExpiresAt)
		if err != nil {
			t.Fatalf("Failed to query user segment: %v", err)
		}

		if !dbExpiresAt.Equal(newExpiresAt) {
			t.Errorf("Expected expires_at to be %v, got %v", newExpiresAt, dbExpiresAt)
		}
	})
}

func TestSegmentRepository_GetUserSegments(t *testing.T) {
	setupDB(t)
	repo := NewSegmentRepository(db)
	ctx := context.Background()

	userId := 1
	segmentId1, _ := repo.CreateSegment(ctx, "AVITO_VOICE_MESSAGES", 0)
	segmentId2, _ := repo.CreateSegment(ctx, "AVITO_PERFORMANCE_VAS", 0)
	_,err := db.ExecContext(ctx,"INSERT INTO users (id) VALUES ($1)",userId)
	if err != nil {
		t.Fatal(err)
	}
	expiresAt := time.Now().Add(time.Hour)
	_ = repo.CreateUserSegment(ctx, userId, segmentId1, expiresAt)
	_ = repo.CreateUserSegment(ctx, userId, segmentId2, expiresAt)
	_ = repo.DeleteUserSegment(ctx,userId,segmentId2)

	t.Run("Success", func(t *testing.T) {
		segments, err := repo.GetUserSegments(ctx, userId)
		if err != nil {
			t.Fatalf("Failed to get user segments: %v", err)
		}

		if len(segments) != 1 {
			t.Errorf("Expected 1 segments, got %d", len(segments))
		}

		expectedSegment := models.GetUserSegmentsResponse{
			SegmentId: segmentId1,
			Slug:      "AVITO_VOICE_MESSAGES",
		}

		if segments[0].SegmentId != expectedSegment.SegmentId {
			t.Errorf("Expected segment ID %d, got %d", expectedSegment.SegmentId, segments[0].SegmentId)
		}
		
		if segments[0].Slug != expectedSegment.Slug {
			t.Errorf("Expected segment slug %s, got %s", expectedSegment.Slug, segments[0].Slug)
		}

		if segments[0].DeletedAt != nil {
			t.Errorf("Expected DeletedAt to be nil, got %v", segments[0].DeletedAt)
		}п
	})

	t.Run("No segments", func(t *testing.T) {
		segments, err := repo.GetUserSegments(ctx, 9999)
		if err != nil {
			t.Fatalf("Failed to get user segments: %v", err)
		}

		if len(segments) != 0 {
			t.Errorf("Expected 0 segments, got %d", len(segments))
		}
	})
}