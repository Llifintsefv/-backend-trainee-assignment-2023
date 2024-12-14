package user

import (
	"backend-trainee-assignment-2023/internal/core/interfaces"
	"context"
	"database/sql"
)


type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) interfaces.UserRepository {
	return &userRepository{db: db}
}



func(r *userRepository) CreateUser(ctx context.Context,id int) error {
	_,err := r.db.ExecContext(ctx,"INSERT INTO users (id) VALUES ($1)",id)
	if err != nil {
		return err
	}
	return nil
}