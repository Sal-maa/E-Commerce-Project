package service

import (
	_"fmt"
	_"time"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/repo"
	_"golang.org/x/crypto/bcrypt"
)

type ProductService interface{
	GetAllProductsService() ([]entity.Product, error)
	GetProductByIdService(id int) (entity.Product, error)
}

type productService struct {
	repository repo.ProductRepository
}

func NewProductService(repository repo.ProductRepository) *productService{
	return &productService{repository}
}

func (s *productService) GetAllProductsService() ([]entity.Product, error) {
	products, err := s.repository.GetAllProducts()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *productService) GetProductByIdService(id int) (entity.Product, error){
	product, err := s.repository.GetProductById(id)
	if err != nil {
		return product, err
	}
	return product, nil
}