package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/helper"
	"github.com/Sal-maa/E-Commerce-Project/service"
	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	cartService service.CartService
}

func NewCartHandler(cartService service.CartService) *cartHandler {
	return &cartHandler{cartService}
}

func (h *cartHandler) CreateCartController(c echo.Context) error {
	cart := entity.Cart{}
	userId := c.Get("currentUser").(entity.User)
	if cart.UserId != userId.Id {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("you dont have permission"))
	}

	cartCreate := entity.CreateCartRequest{}
	if err := c.Bind(&cartCreate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}

	_, err := h.cartService.CreateCartService(cartCreate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to insert data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithoutDataResponses("success insert data"))
}

func (h *cartHandler) GetAllCartsController(c echo.Context) error {
	cart := entity.Cart{}
	userId := c.Get("currentUser").(entity.User)
	if cart.UserId != userId.Id {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("you dont have permission"))
	}
	carts, err := h.cartService.GetAllCartsService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerError("failed to fetch data"))
	}

	data := []entity.CartResponse{}
	for i := 0; i < len(carts); i++ {
		formatRes := entity.FormatCartResponse(carts[i])
		data = append(data, formatRes)
	}

	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", data))
}

func (h *cartHandler) DeleteCartController(c echo.Context) error {
	cart := entity.Cart{}
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed convert id"))
	}

	userId := c.Get("currentUser").(entity.User)
	if cart.UserId != userId.Id {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("you dont have permission"))
	}

	_, err1 := h.cartService.DeleteCartService(idParam)

	if err1 != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed delete data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithoutDataResponses("Success delete data"))
}
