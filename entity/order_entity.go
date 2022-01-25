package entity

import "time"

type Address struct {
	Id     int    `json:"id" form:"id"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    int    `json:"zip"`
}

type CreditCard struct {
	Id     int    `json:"id" form:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Number string `json:"number"`
	CVV    int    `json:"cvv"`
}

type OrderDetails struct {
	Id      int `json:"id" form:"id"`
	CartId  int `json:"cart_id" form:"cart_id"`
	OrderId int `json:"order_id" form:"order_id"`
}

type Order struct {
	Id          int       `json:"id" form:"id"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" form:"deleted_at"`
	User        User      `json:"user"`
	Cart        []int
	Address     Address    `json:"address"`
	CreditCard  CreditCard `json:"credit_card"`
	StatusOrder string     `json:"status_order"`
	OrderDate   time.Time  `json:"order_date"`
	Total       int        `json:"total"`
}

type CreateOrderRequest struct {
	User       User
	CartId     []int      `json:"cart_id"`
	Address    Address    `json:"address"`
	CreditCard CreditCard `json:"credit_card"`
	OrderDate  time.Time  `json:"order_date"`
	Total      int        `json:"total"`
}

type CreateOrderDetailRequest struct {
	CartId  int `json:"cart_id" form:"cart_id"`
	OrderId int `json:"order_id" form:"order_id"`
}

type EditOrderRequest struct {
	StatusOrder string `json:"status_order"`
}

type OrderResponse struct {
	Id			int				  `json:"id"`
	User        UserOrderResponse `json:"user"`
	Cart        []int 			  `json:"cart"`
	Address     Address           `json:"address"`
	StatusOrder string            `json:"status_order"`
	OrderDate   time.Time         `json:"order_date"`
	Total       int               `json:"total"`
}

func FormatOrderResponse(order Order) OrderResponse {
	return OrderResponse{
		Id:	order.Id,
		User: UserOrderResponse{
			Id:       order.User.Id,
			Username: order.User.Username,
		},
		Cart: 		 order.Cart,
		Address:     Address{
			Id: 	order.Address.Id,
			Street: order.Address.Street,
			City: order.Address.City,
			State: order.Address.State,
			Zip: order.Address.Zip,
		},
		StatusOrder: order.StatusOrder,
		OrderDate:   order.OrderDate,
		Total:       order.Total,
	}
}
