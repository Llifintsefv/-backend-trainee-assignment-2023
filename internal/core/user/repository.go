package user

import (
	"backend-trainee-assignment-2023/internal/core/interfaces"
	"database/sql"
)


type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) interfaces.UserRepository {
	return &userRepository{db: db}
}