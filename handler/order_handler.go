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

func (h *orderHandler) UpdateOrderController(c echo.Context) error {
	updatedOrder := entity.EditOrderRequest{}
	if err := c.Bind(&updatedOrder); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}

	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed convert id"))
	}

	_, err = h.orderService.UpdateOrderService(idParam, updatedOrder)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to update data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithoutDataResponses("success insert data"))
}

func (h *orderHandler) GetOrderController(c echo.Context) error {
	userId := c.Get("currentUser").(entity.User)

	orders, err := h.orderService.GetOrderService(userId.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerError("failed to fetch data"))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerError("failed to fetch data"))
	}

	data := []entity.OrderResponse{}
	for i := 0; i < len(orders); i++ {
		formatRes := entity.FormatOrderResponse(orders[i])
		data = append(data, formatRes)
	}

	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", data))
}