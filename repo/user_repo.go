package repo

import (
	"database/sql"
)

type UserRepository interface {
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}
