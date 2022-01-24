package entity

import "time"

type Cart struct {
	Id        int       `json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
	Product   Product
	User      User
	Qty       int `json:"qty" form:"qty"`
	Subtotal  int `json:"subtotal" form:"subtotal"`
}

type CreateCartRequest struct {
	ProductId int `json:"product_id"`
	User      User
	Qty       int `json:"qty" form:"qty"`
	Subtotal  int `json:"subtotal" form:"subtotal"`
}

type EditCartRequest struct {
	ProductId int `json:"product_id"`
	User      User
	Qty       int `json:"qty" form:"qty"`
	Subtotal  int `json:"subtotal" form:"subtotal"`
}

type CartResponse struct {
	User     UserProductResponse `json:"user"`
	Product  ProductCartResponse `json:"product"`
	Qty      int                 `json:"qty" form:"qty"`
	Subtotal int                 `json:"subtotal" form:"subtotal"`
}

func FormatCartResponse(cart Cart) CartResponse {
	return CartResponse{
		User: UserProductResponse{
			Id:       cart.Product.User.Id,
			Username: cart.Product.User.Username,
		},
		Product: ProductCartResponse{
			Id:     cart.Product.Id,
			Name:   cart.Product.Name,
			Gambar: cart.Product.Gambar,
			Harga:  cart.Product.Harga,
		},
		Qty:      cart.Qty,
		Subtotal: cart.Subtotal,
	}
}
