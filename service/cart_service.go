package service

import (
	"time"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/repo"
)

type CartService interface {
	CreateCartService(cartCreate entity.CreateCartRequest) (entity.Cart, error)
	GetAllCartsService() ([]entity.Cart, error)
	GetCartByIdService(id int) (entity.Cart, error)
	DeleteCartService(id int) (entity.Cart, error)
	UpdateCartService(id int, cart entity.EditCartRequest) (entity.Cart, error)
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

func (s *cartService) GetAllCartsService() ([]entity.Cart, error) {
	carts, err := s.repository.GetAllCarts()
	if err != nil {
		return carts, err
	}
	return carts, nil
}

func (s *cartService) GetCartByIdService(id int) (entity.Cart, error) {
	cart, err := s.repository.GetCartById(id)
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (s *cartService) DeleteCartService(id int) (entity.Cart, error) {
	cartID, err := s.GetCartByIdService(id)
	if err != nil {
		return cartID, err
	}

	cartID.DeletedAt = time.Now()
	deleteCart, err := s.repository.DeleteCart(cartID)

	return deleteCart, err
}

func (s *cartService) UpdateCartService(id int, cart entity.EditCartRequest) (entity.Cart, error){
	cartId, err := s.GetCartByIdService(id)
	if err != nil {
		return cartId, err
	}
	cartId.UpdatedAt = time.Now()
	cartId.Product_Id = cart.Product_Id
	cartId.Qty = cart.Qty
	cartId.Subtotal = cart.Subtotal
	cartId.UserId = cart.UserId

	updateCart, err := s.repository.UpdateCart(cartId)
	return	updateCart, err
}