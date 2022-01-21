package repo

import (
	"database/sql"

	"github.com/Sal-maa/E-Commerce-Project/entity"
)

type CartRepository interface {
	CreateCart(cart entity.Cart) (entity.Cart, error)
}

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *cartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) CreateCart(cart entity.Cart) (entity.Cart, error) {
	_, err := r.db.Exec("INSERT INTO carts(created_at, updated_at, product_id,qty,subtotal) VALUES(?,?,?,?,?)", cart.CreatedAt, cart.UpdatedAt, cart.Product_Id, cart.Qty, cart.Subtotal)
	return cart, err
}
