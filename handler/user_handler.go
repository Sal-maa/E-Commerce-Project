package handler

import (
	"fmt"
	"net/http"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/helper"
	"github.com/Sal-maa/E-Commerce-Project/middleware"
	"github.com/Sal-maa/E-Commerce-Project/service"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	authService middleware.JWTService
	userService service.UserService
}

func NewUserHandler(authService middleware.JWTService, userService service.UserService) *userHandler {
	return &userHandler{authService, userService}
}

func (h *userHandler) CreateUserController(c echo.Context) error {
	userCreate := entity.CreateUserRequest{}
	if err := c.Bind(&userCreate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}
	_, err := h.userService.CreateUserService(userCreate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to insert data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithoutDataResponses("success insert data"))
}

func (h *userHandler) AuthController(c echo.Context) error {
	login := entity.LoginUserRequest{}
	if err := c.Bind(&login); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("input must be string"))
	}

	user, err := h.userService.LoginUserService(login)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("Wrong combination username and password"))
	}

	// membuat token
	token, err := h.authService.GenerateToken(user.Name)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, helper.FailedResponses("cannot create token"))
	}
	saveToken, err := h.userService.SaveTokenService(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponses("cannot save token"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponses("login success", saveToken))
}
