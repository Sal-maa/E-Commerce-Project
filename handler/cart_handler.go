package handler

import (
	"fmt"
	"net/http"

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
