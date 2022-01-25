package service

import (
	"fmt"
	_ "fmt"
	"time"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/repo"
)

type OrderService interface {
	CreateOrderService(orderCreate entity.CreateOrderRequest) (entity.Order, error)
	UpdateOrderService(id int, updatedOrder entity.EditOrderRequest) (entity.Order, error)
	GetOrderService(id int) ([]entity.Order, error)
	GetOrderByIdService(userId, orderId int) (entity.Order, error)
}

type orderService struct {
	repository repo.OrderRepository
}

func NewOrderService(repository repo.OrderRepository) *orderService {
	return &orderService{repository}
}

func (s *orderService) CreateOrderService(orderCreate entity.CreateOrderRequest) (entity.Order, error) {
	order := entity.Order{}

	order.Address.Street = orderCreate.Address.Street
	order.Address.City = orderCreate.Address.City
	order.Address.State = orderCreate.Address.State
	order.Address.Zip = orderCreate.Address.Zip
	createAddress, errAddress := s.repository.CreateAddress(order)
	if errAddress != nil {
		return order, errAddress
	}

	order.CreditCard.Type = orderCreate.CreditCard.Type
	order.CreditCard.Name = orderCreate.CreditCard.Name
	order.CreditCard.Number = orderCreate.CreditCard.Number
	order.CreditCard.CVV = orderCreate.CreditCard.CVV

	createPayment, errPayment := s.repository.CreatePayment(order)
	if errPayment != nil {
		return order, errPayment
	}

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.User = orderCreate.User
	order.Address.Id = createAddress.Id
	order.CreditCard.Id = createPayment.Id
	order.StatusOrder = "Processed"
	order.OrderDate = time.Now()
	order.Total = orderCreate.Total
	order.Cart = orderCreate.CartId

	createOrder, errOrder := s.repository.CreateOrder(order)
	if errOrder != nil {
		return order, errOrder
	}
	fmt.Println(createOrder)
	orderDetail := entity.CreateOrderDetailRequest{}
	orderDetail.OrderId = createOrder
	for _, v := range orderCreate.CartId {
		orderDetail.CartId = v
		_, errOrder := s.repository.CreateOrderDetail(orderDetail)
		if errOrder != nil {
			return order, errOrder
		}
	}
	return order, errOrder
}

func (s *orderService) UpdateOrderService(id int, updatedOrder entity.EditOrderRequest) (entity.Order, error) {
	order := entity.Order{}
	order.Id = id
	order.UpdatedAt = time.Now()
	order.StatusOrder = updatedOrder.StatusOrder

	orderUpdated, err := s.repository.UpdateOrder(order)
	if err != nil {
		return order, err
	}
	return orderUpdated, nil
}

func (s *orderService) GetOrderService(id int) ([]entity.Order, error){
	orders, err := s.repository.GetOrder(id)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *orderService) GetOrderByIdService(userId, orderId int) (entity.Order, error){
	order, err := s.repository.GetOrderById(userId, orderId)
	if err != nil {
		return order, err
	}
	return order, nil
}