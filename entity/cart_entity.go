package entity

import "time"

type Cart struct {
	Id         int       `json:"id" form:"id"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at" form:"deleted_at"`
	Product_Id int       `json:"product_id" form:"product_id"`
	Qty        int       `json:"qty" form:"qty"`
	Subtotal   int       `json:"subtotal" form:"subtotal"`
}

type CreateCartRequest struct {
	Product_Id int `json:"product_id" form:"product_id"`
	Qty        int `json:"qty" form:"qty"`
	Subtotal   int `json:"subtotal" form:"subtotal"`
}

type EditCartRequest struct {
	Product_Id int `json:"product_id" form:"product_id"`
	Qty        int `json:"qty" form:"qty"`
	Subtotal   int `json:"subtotal" form:"subtotal"`
}

type CartResponse struct {
	Product_Id int `json:"product_id" form:"product_id"`
	Qty        int `json:"qty" form:"qty"`
	Subtotal   int `json:"subtotal" form:"subtotal"`
}

func FormatCartResponse(cart Cart) CartResponse {
	return CartResponse{
		Product_Id: cart.Product_Id,
		Qty:        cart.Qty,
		Subtotal:   cart.Subtotal,
	}
}
