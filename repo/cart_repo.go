package repo

import (
	"database/sql"
	"fmt"

	"github.com/Sal-maa/E-Commerce-Project/entity"
)

type CartRepository interface {
	CreateCart(cart entity.Cart) (entity.Cart, error)
	GetAllCarts(userId entity.User) ([]entity.Cart, error)
	GetCartById(id int) (entity.Cart, error)
	DeleteCart(cart entity.Cart) (entity.Cart, error)
	UpdateCart(cart entity.Cart) (entity.Cart, error)
	ShowCart(updateCart entity.Cart) (entity.Cart, error)
}

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *cartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) CreateCart(cart entity.Cart) (entity.Cart, error) {
	_, err := r.db.Exec(`INSERT INTO carts(created_at, updated_at, product_id, user_id, qty, subtotal)
						VALUES(?,?,?,?,?,?)`, cart.CreatedAt, cart.UpdatedAt, cart.Product.Id, cart.User.Id, cart.Qty, cart.Subtotal)
	return cart, err
}

func (r *cartRepository) GetAllCarts(userId entity.User) ([]entity.Cart, error) {
	carts := []entity.Cart{}
	result, err := r.db.Query(`SELECT u.username, c.id, c.product_id, p.user_id, p.name, p.gambar, p.harga, c.qty, c.subtotal 
								FROM carts c JOIN products p ON p.id=c.product_id
								JOIN users u ON u.id = p.user_id WHERE c.user_id=?`, userId.Id)
	if err != nil {
		fmt.Println(err)
		return carts, fmt.Errorf("failed in query")
	}

	defer result.Close()

	for result.Next() {
		cart := entity.Cart{}
		err := result.Scan(&cart.Product.User.Username, &cart.Id, &cart.Product.Id, &cart.Product.User.Id, &cart.Product.Name, &cart.Product.Gambar, &cart.Product.Harga, &cart.Qty, &cart.Subtotal)
		if err != nil {
			return carts, fmt.Errorf("failed to scan")
		}
		carts = append(carts, cart)
	}
	return carts, err
}

func (r *cartRepository) GetCartById(id int) (entity.Cart, error) {
	cart := entity.Cart{}
	result, err := r.db.Query("SELECT id, user_id, product_id, qty, subtotal FROM carts WHERE id = ?", id)
	if err != nil {
		return cart, fmt.Errorf("failed in query")
	}
	defer result.Close()
	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}
	errScan := result.Scan(&cart.Id, &cart.User.Id, &cart.Product.Id, &cart.Qty, &cart.Subtotal)
	if errScan != nil {
		fmt.Println(errScan)
		return cart, fmt.Errorf("failed to read data")
	}
	if id == cart.Id {
		return cart, nil
	}
	return cart, fmt.Errorf("cart not found")
}

func (r *cartRepository) UpdateCart(cart entity.Cart) (entity.Cart, error) {
	_, err := r.db.Exec("UPDATE carts SET updated_at = ?, qty = ?, subtotal = ? WHERE id = ?", cart.UpdatedAt, cart.Qty, cart.Subtotal, cart.Id)
	if err != nil {
		fmt.Println("update error:", err)
	}
	return cart, err
}

func (r *cartRepository) ShowCart(updateCart entity.Cart) (entity.Cart, error) {
	result, err := r.db.Query(`SELECT u.username, c.id, c.product_id, p.user_id, p.name, p.gambar, p.harga, c.qty, c.subtotal
								FROM carts c JOIN products p ON p.id=c.product_id
								JOIN users u ON u.id = p.user_id WHERE c.id=?`, updateCart.Id)
	if err != nil {
		return updateCart, fmt.Errorf("failed in query")
	}

	defer result.Close()

	if isExist := result.Next(); isExist {
		fmt.Println(isExist)
		fmt.Println("data is exist")
	}

	errScan := result.Scan(&updateCart.Product.User.Username, &updateCart.Id, &updateCart.Product.Id, &updateCart.Product.User.Id, &updateCart.Product.Name, &updateCart.Product.Gambar, &updateCart.Product.Harga, &updateCart.Qty, &updateCart.Subtotal)
	if errScan != nil {
		fmt.Println(errScan)
		return updateCart, fmt.Errorf("failed to read data")
	}
	return updateCart, err
}

// func (r *cartRepository) DeleteCart(cart entity.Cart) (entity.Cart, error) {
// 	_, err := r.db.Exec("UPDATE carts SET deleted_at=? WHERE id=?", cart.DeletedAt, cart.Id)
// 	return cart, err
// }

func (r *cartRepository) DeleteCart(cart entity.Cart) (entity.Cart, error) {
	_, err := r.db.Exec("DELETE FROM carts WHERE id=?", cart.Id)
	return cart, err
}
