package repo

import (
	"database/sql"
	"fmt"

	"github.com/Sal-maa/E-Commerce-Project/entity"
)

type CartRepository interface {
	CreateCart(cart entity.Cart) (entity.Cart, error)
	GetAllCarts() ([]entity.Cart, error)
	GetCartById(id int) (entity.Cart, error)
	DeleteCart(cart entity.Cart) (entity.Cart, error)
}

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *cartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) CreateCart(cart entity.Cart) (entity.Cart, error) {
	_, err := r.db.Exec("INSERT INTO carts(created_at, updated_at, product_id, qty, subtotal) VALUES(?,?,?,?,?)", cart.CreatedAt, cart.UpdatedAt, cart.Product_Id, cart.Qty, cart.Subtotal)
	return cart, err
}

func (r *cartRepository) GetAllCarts() ([]entity.Cart, error) {
	carts := []entity.Cart{}

	result, err := r.db.Query("SELECT id, product_id, qty, subtotal FROM carts")
	if err != nil {
		fmt.Println(err)
		return carts, fmt.Errorf("failed in query")
	}

	defer result.Close()

	for result.Next() {
		cart := entity.Cart{}
		err := result.Scan(&cart.Id, &cart.Product_Id, &cart.Qty, &cart.Subtotal)
		if err != nil {
			return carts, fmt.Errorf("failed to scan")
		}
		carts = append(carts, cart)
	}
	return carts, err
}

func (r *cartRepository) GetCartById(id int) (entity.Cart, error) {
	cart := entity.Cart{}
	result, err := r.db.Query("SELECT id, product_id, qty, subtotal FROM carts WHERE id = ?", id)
	if err != nil {
		return cart, fmt.Errorf("failed in query")
	}
	defer result.Close()
	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}
	errScan := result.Scan(&cart.Id, &cart.Product_Id, &cart.Qty, &cart.Subtotal)
	fmt.Println(errScan)
	if errScan != nil {

		return cart, fmt.Errorf("failed to read data")
	}
	if id == cart.Id {
		return cart, nil
	}
	return cart, fmt.Errorf("cart not found")
}

func (r *cartRepository) DeleteCart(cart entity.Cart) (entity.Cart, error) {
	_, err := r.db.Exec("UPDATE carts SET deleted_at=? WHERE id=?", cart.DeletedAt, cart.Id)
	return cart, err
}
