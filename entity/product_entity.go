package entity

import "time"

type Product struct {
	Id         int       `json:"id" form:"id"`
	UpdatedAt  time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at" form:"deleted_at"`
	UserId     int       `json:"user_id" form:"user_id"`
	CategoryId int       `json:"category_id" form:"category_id"`
	Name       string    `json:"name" form:"name"`
	Deskripsi  string    `json:"deskripsi" form:"deskripsi"`
	Gambar     string    `json:"gambar" form:"gambar"`
	Harga      int       `json:"harga" form:"harga"`
	Stock      int       `json:"stock" form:"stock"`
}

type CreateProduct struct {
	UserId     int    `json:"user_id" form:"user_id"`
	CategoryId int    `json:"category_id" form:"category_id"`
	Name       string `json:"name" form:"name"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
	Gambar     string `json:"gambar" form:"gambar"`
	Harga      int    `json:"harga" form:"harga"`
	Stock      int    `json:"stock" form:"stock"`
}

type EditProduct struct {
	UserId     int    `json:"user_id" form:"user_id"`
	CategoryId int    `json:"category_id" form:"category_id"`
	Name       string `json:"name" form:"name"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
	Gambar     string `json:"gambar" form:"gambar"`
	Harga      int    `json:"harga" form:"harga"`
	Stock      int    `json:"stock" form:"stock"`
}

type ProductResponse struct {
	Id         int    `json:"id" form:"id"`
	UserId     int    `json:"user_id" form:"user_id"`
	CategoryId int    `json:"category_id" form:"category_id"`
	Name       string `json:"name" form:"name"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
	Gambar     string `json:"gambar" form:"gambar"`
	Harga      int    `json:"harga" form:"harga"`
	Stock      int    `json:"stock" form:"stock"`
}

func FormatProductResponse(product Product) ProductResponse {
	return ProductResponse{
		Id:         product.Id,
		UserId:     product.UserId,
		CategoryId: product.CategoryId,
		Name:       product.Name,
		Deskripsi:  product.Deskripsi,
		Gambar:     product.Gambar,
		Harga:      product.Harga,
		Stock:      product.Stock,
	}
}
