package repo

import (
	"database/sql"
	"fmt"

	"github.com/Sal-maa/E-Commerce-Project/entity"
)

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{db}
}



func (r *productRepository) GetAllProducts() ([]entity.Product, error) {
	products := []entity.Product{}

	result, err := r.db.Query("SELECT id, name, deskripsi, gambar, harga, stock, category_id, user_id FROM products")
	if err != nil {
		fmt.Println(err)
		return products, fmt.Errorf("failed to scan")
	}

	defer result.Close()

	for result.Next() {
		product := entity.Product{}
		err := result.Scan(&product.Id, &product.Name, &product.Deskripsi, &product.Gambar, &product.Harga, &product.Stock, &product.CategoryId, &product.UserId)
		if err != nil {
			return products, fmt.Errorf("failed to scan")
		}
		products = append(products, product)
	}
	return products, err
}