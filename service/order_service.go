package service

import (
	"time"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/repo"
)

type OrderService interface {
	CreateOrderService(orderCreate entity.CreateOrderRequest) (entity.Order, error)
	UpdateOrderService(orderId int, editOrder entity.EditOrderRequest) (entity.Order, error)
}

type orderService struct {
	repository repo.OrderRepository
}

func NewOrderService(repository repo.OrderRepository) *orderService {
	return &orderService{repository}
}

func (s *orderService) CreateOrderService(orderCreate entity.CreateOrderRequest) (entity.Order, error) {
	order := entity.Order{}

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.Cart.Id = orderCreate.CartId
	order.StatusOrder = "Processed"
	order.OrderDate = time.Now()
	order.Total = orderCreate.Total
	createOrder1, err1 := s.repository.CreateOrder1(order)
	if err1 != nil {
		return createOrder1, err1
	}
	order.Address.Street = orderCreate.Address.Street
	order.Address.City = orderCreate.Address.City
	order.Address.State = orderCreate.Address.State
	order.Address.Zip = orderCreate.Address.Zip
	createOrder2, err2 := s.repository.CreateOrder2(order)
	if err2 != nil {
		return createOrder2, err2
	}
	order.CreditCard.Type = orderCreate.CreditCard.Type
	order.CreditCard.Name = orderCreate.CreditCard.Name
	order.CreditCard.Number = orderCreate.CreditCard.Number
	order.CreditCard.CVV = orderCreate.CreditCard.CVV
	createOrder3, err3 := s.repository.CreateOrder3(order)

	return createOrder3, err3
}

func (s *orderService) UpdateOrderService(orderId int, editOrder entity.EditOrderRequest) (entity.Order, error) {
	order := entity.Order{}
	order.Id = orderId
	order.UpdatedAt = time.Now()
	order.StatusOrder = editOrder.StatusOrder

	orderUpdate, err := s.repository.UpdateOrder(order)
	if err != nil{
		return orderUpdate, err
	}
	return orderUpdate, nil	
}