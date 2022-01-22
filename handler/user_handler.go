package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

	if strings.Contains(userCreate.Username, " ") {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("username can't contain space"))
	}
	_, err1 := h.userService.CheckUserName(userCreate)
	if err1 != nil {
		fmt.Println(err1)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("username already exist"))
	}

	_, err := h.userService.CreateUserService(userCreate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to insert data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithoutDataResponses("success insert data"))
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
	token, err := h.authService.GenerateToken(user.Username)
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

func (h *userHandler) GetUserController(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("data not found"))
	}

	user, err := h.userService.GetUserByIdService(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerError("failed to fetch data"))
	}

	formatRes := entity.FormatUserResponse(user)
	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", formatRes))
}

func (h *userHandler) DeleteUserController(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed convert id"))
	}

	user, err1 := h.userService.DeleteUserService(idParam)
	if err1 != nil {
		fmt.Println(err1)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed delete data"))
	}
	formatRes := entity.FormatUserResponse(user)
	return c.JSON(http.StatusOK, helper.SuccessResponses("Success delete data", formatRes))
}

func (h *userHandler) UpdateUserController(c echo.Context) error {
	userUpdate := entity.EditUserRequest{}
	if err := c.Bind(&userUpdate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}

	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed convert id"))
	}

	_, err = h.userService.UpdateUserService(idParam, userUpdate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to update data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithoutDataResponses("success update data"))
}
