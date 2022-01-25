package service

import (
	"fmt"
	_ "fmt"
	"time"
	_ "time"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/repo"
	_ "golang.org/x/crypto/bcrypt"
)

type ProductService interface {
	GetAllProductsService() ([]entity.Product, error)
	GetAllUserProductsService(userId int) ([]entity.Product, error)
	GetProductByIdService(id int) (entity.Product, error)
	CreateProductService(userId int, product entity.CreateProduct) (entity.Product, error)
	UpdateProductService(id int, productUpdate entity.EditProduct) (entity.Product, error)
	DeleteProductService(id, currentUser int) (entity.Product, error)
}

type productService struct {
	repository repo.ProductRepository
}

func NewProductService(repository repo.ProductRepository) *productService {
	return &productService{repository}
}

func (s *productService) GetAllProductsService() ([]entity.Product, error) {
	products, err := s.repository.GetAllProducts()
	return products, err
}

func (s *productService) GetAllUserProductsService(userId int) ([]entity.Product, error) {
	products, err := s.repository.GetAllUserProducts(userId)
	return products, err
}

func (s *productService) GetProductByIdService(id int) (entity.Product, error) {
	product, err := s.repository.GetProductById(id)
	return product, err
}

func (s *productService) UpdateProductService(id int, productUpdate entity.EditProduct) (entity.Product, error) {
	product, err := s.repository.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return product, err
	}

	if productUpdate.User.Id != product.User.Id {
		return product, fmt.Errorf("you dont have permission")
	}

	product.UpdatedAt = time.Now()
	product.Name = productUpdate.Name
	product.Deskripsi = productUpdate.Deskripsi
	product.Gambar = productUpdate.Gambar
	product.Harga = productUpdate.Harga
	product.Stock = productUpdate.Stock
	product.CategoryId = productUpdate.CategoryId

	updateProduct, err := s.repository.UpdateProduct(product)
	return updateProduct, err
}

func (s *productService) DeleteProductService(id, currentUser int) (entity.Product, error) {
	productID, err := s.GetProductByIdService(id)
	if err != nil {
		return productID, err
	}

	if productID.User.Id != currentUser {
		return productID, fmt.Errorf("you dont have permission")
	}

	productID.DeletedAt = time.Now()
	deleteProduct, err := s.repository.DeleteProduct(productID)

	return deleteProduct, err
}

func (s *productService) CreateProductService(userId int, product entity.CreateProduct) (entity.Product, error) {
	produk := entity.Product{}
	produk.CreatedAt = time.Now()
	produk.UpdatedAt = time.Now()
	produk.Name = product.Name
	produk.Deskripsi = product.Deskripsi
	produk.Gambar = product.Gambar
	produk.Harga = product.Harga
	produk.Stock = product.Stock
	produk.CategoryId = product.CategoryId
	produk.User.Id = userId

	storedProduct, err := s.repository.CreateProduct(userId, produk)
	return storedProduct, err
}
