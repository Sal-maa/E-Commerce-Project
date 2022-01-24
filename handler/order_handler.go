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

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *orderHandler {
	return &orderHandler{orderService}
}

func (h *orderHandler) CreateOrderController(c echo.Context) error {
	orderCreate := entity.CreateOrderRequest{}

	userId := c.Get("currentUser").(entity.User)
	orderCreate.User = userId

	if err := c.Bind(&orderCreate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}

	_, err := h.orderService.CreateOrderService(orderCreate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to insert data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithoutDataResponses("success insert data"))

}

func (h *orderHandler) UpdateOrderController(c echo.Context) error{
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed convert id"))
	}

	orderUpdate := entity.EditOrderRequest{}
	if err := c.Bind(&orderUpdate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}

	_, err = h.orderService.UpdateOrderService(idParam, orderUpdate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to update data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithoutDataResponses("success insert data"))
}