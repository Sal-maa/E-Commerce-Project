package handler

import "github.com/Sal-maa/E-Commerce-Project/service"

type userHandler struct {
	authService middleware.JWTService
	userService service.UserService
}

func NewUserHandler(authService middleware.JWTService, userService service.UserService) *userHandler {
	return &userHandler{authService, userService}
}
