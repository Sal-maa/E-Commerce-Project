package repo

import (
	"database/sql"
	"fmt"

	"github.com/Sal-maa/E-Commerce-Project/entity"
)

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductById(id int) (entity.Product, error)
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

func (r *productRepository) GetProductById(id int) (entity.Product, error){
	product := entity.Product{}
	result, err := r.db.Query("SELECT id, name, deskripsi, gambar, harga, stock, category_id, user_id FROM products WHERE id = ?", id)
	if err != nil {
		return product, fmt.Errorf("failed in query")
	}
	defer result.Close()
	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}
	errScan := result.Scan(&product.Id, &product.Name, &product.Deskripsi, &product.Gambar, &product.Harga, &product.Stock, &product.CategoryId, &product.UserId)
	fmt.Println(errScan)
	if errScan != nil {

		return product, fmt.Errorf("failed to read data")
	}
	if id == product.Id {
		return product, nil
	}
	return product, fmt.Errorf("product not found")
}