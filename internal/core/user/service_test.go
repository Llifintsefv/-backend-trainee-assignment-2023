package user

import (
	"backend-trainee-assignment-2023/internal/core/models"
	"context"
	"errors"
	"testing"
)

type MockUserRepo struct {
	CreateUserFunc func(ctx context.Context,userId int) error
	UserExistsFunc func(userId int) (bool, error)
}

func (m *MockUserRepo) CreateUser(ctx context.Context,userId int) error {
	return m.CreateUserFunc(ctx,userId)
}

func (m *MockUserRepo) UserExists(userId int) (bool, error) {
	return m.UserExistsFunc(userId)
}


func TestCreateUser_Success(t *testing.T){
	userId := 1
	ctx := context.Background()

	user := models.User{Id: userId}

	mockUserRepo := &MockUserRepo{
		CreateUserFunc: func(ctx context.Context, userId int) error {
			if userId != 1 {
				t.Errorf("expecred userId 1, got %d",userId)
			}
			return nil
		},
	}

	service := NewUserService(mockUserRepo,nil)

	err := service.CreateUser(ctx,user)
	if err != nil {
		t.Errorf("Expected nil,got %d",err)
	}
}


func TestCreateUser_Error(t *testing.T) {
	userId := 1
	user := models.User{Id: userId}
	ctx := context.Background()

	MockUserRepo := &MockUserRepo{
		CreateUserFunc: func(ctx context.Context, userId int) error {
			return errors.New("database error")
		},
	}

	service := NewUserService(MockUserRepo,nil)

	err := service.CreateUser(ctx,user)

	if err == nil {
		t.Errorf("expected err, got nil")
	}

	if err.Error() != "database error" {
		t.Errorf("Expected error message 'database error', got '%s'", err.Error())

	}
}