package entity

type UserProductResponse struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
}

type ProductCartResponse struct {
	Id     int    `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Gambar string `json:"gambar" form:"gambar"`
	Harga  int    `json:"harga" form:"harga"`
}
