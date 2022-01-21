package repo

import (
	"database/sql"
	_"fmt"

	"github.com/Sal-maa/E-Commerce-Project/entity"
)

type ProductRepository interface {
	GetProducts() (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}