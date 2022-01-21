package service

import (
	"time"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/repo"
)

type CartService interface {
	CreateCartService(cartCreate entity.CreateCartRequest) (entity.Cart, error)
}

type cartService struct {
	repository repo.CartRepository
}

func NewCartService(repository repo.CartRepository) *cartService {
	return &cartService{repository}
}

func (s *cartService) CreateCartService(cartCreate entity.CreateCartRequest) (entity.Cart, error) {
	cart := entity.Cart{}
	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()
	cart.Product_Id = cartCreate.Product_Id
	cart.Qty = cartCreate.Qty
	cart.Subtotal = cartCreate.Subtotal

	createCart, err := s.repository.CreateCart(cart)
	return createCart, err
}
