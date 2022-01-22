package entity

type UserProductResponse struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
}
