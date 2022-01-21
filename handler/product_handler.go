package handler

import (
	"fmt"
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/helper"
	_ "github.com/Sal-maa/E-Commerce-Project/middleware"
	"github.com/Sal-maa/E-Commerce-Project/service"
	"github.com/labstack/echo/v4"
)

type productHandler struct {
	//authService middleware.JWTService
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) GetAllProductsController(c echo.Context) error {
	products, err := h.productService.GetAllProductsService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerError("failed to fetch data"))
	}

	data := []entity.ProductResponse{}
	for i := 0; i < len(products); i++ {
		formatRes := entity.FormatProductResponse(products[i])
		data = append(data, formatRes)
	}

	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", data))
}

func (h *productHandler) GetProductByIdController(c echo.Context) error {
	convId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("data not found"))
	}
	product, err := h.productService.GetProductByIdService(convId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to fetch data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", product))
}

func (h *productHandler) UpdateProductController(c echo.Context) error {
	productUpdate := entity.EditProduct{}
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("data not found"))
	}

	if err := c.Bind(&productUpdate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}

	userId := c.Get("currentUser").(entity.User)
	if productUpdate.UserId != userId.Id {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("you dont have permission"))
	}

	productUp, err := h.productService.UpdateProductService(idParam, productUpdate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to update data"))
	}

	formatRes := entity.FormatProductResponse(productUp)
	return c.JSON(http.StatusOK, helper.SuccessResponses("success update data", formatRes))
}

func (h *productHandler) DeleteProductController(c echo.Context) error {
	product := entity.Product{}
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed convert id"))
	}

	userId := c.Get("currentUser").(entity.User)
	if product.UserId != userId.Id {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("you dont have permission"))
	}

	_, err1 := h.productService.DeleteProductService(idParam)

	if err1 != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed delete data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithoutDataResponses("Success delete data"))
}
