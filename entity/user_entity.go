package entity

import "time"

type LoginUserRequest struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

func FormatLoginResponse(token string) LoginUserResponse {
	return LoginUserResponse{
		Token: token,
	}
}

type User struct {
	Id        int       `json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Address   string    `json:"address" form:"address"`
	Phone     string    `json:"phone" form:"phone"`
}

type CreateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
}

type EditUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
}

type UserResponse struct {
	Id      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
	Phone   string `json:"phone" form:"phone"`
}

func FormatUserResponse(user User) UserResponse {
	return UserResponse{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}
}
