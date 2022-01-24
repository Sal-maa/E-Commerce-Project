package service

import (
	"fmt"
	"time"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/repo"
)

type CartService interface {
	CreateCartService(cartCreate entity.CreateCartRequest) (entity.Cart, error)
	GetAllCartsService(userId entity.User) ([]entity.Cart, error)
	GetCartByIdService(id int) (entity.Cart, error)
	DeleteCartService(id, currentUser int) (entity.Cart, error)
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
	cart.Product.Id = cartCreate.ProductId
	cart.User.Id = cartCreate.User.Id
	cart.Qty = cartCreate.Qty
	cart.Subtotal = cartCreate.Subtotal

	createCart, err := s.repository.CreateCart(cart)
	return createCart, err
}

func (s *cartService) GetAllCartsService(userId entity.User) ([]entity.Cart, error) {
	carts, err := s.repository.GetAllCarts(userId)
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

func (s *cartService) DeleteCartService(id, currentUser int) (entity.Cart, error) {
	cartID, err := s.GetCartByIdService(id)
	if err != nil {
		return cartID, err
	}

	if cartID.User.Id != currentUser {
		return cartID, fmt.Errorf("you dont have permission")
	}

	cartID.DeletedAt = time.Now()
	deleteCart, err := s.repository.DeleteCart(cartID)

	return deleteCart, err
}

func (s *cartService) UpdateCartService(id int, cartUpdate entity.EditCartRequest) (entity.Cart, error) {
	cartId, err := s.GetCartByIdService(id)
	if err != nil {
		return cartId, err
	}

	if cartUpdate.User.Id != cartId.User.Id {
		fmt.Println(cartUpdate.User.Id, cartId.User.Id)
		return cartId, fmt.Errorf("you dont have permission")
	}
	cartId.UpdatedAt = time.Now()
	cartId.Qty = cartUpdate.Qty
	cartId.Subtotal = cartUpdate.Subtotal

	updateCart, err := s.repository.UpdateCart(cartId)
	showCart, _ := s.repository.ShowCart(updateCart)
	if err == nil {
		return showCart, err
	}
	return updateCart, err
}
