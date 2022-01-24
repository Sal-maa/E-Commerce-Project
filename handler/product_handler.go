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
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) GetAllProductsController(c echo.Context) error {
	products, err := h.productService.GetAllProductsService()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, helper.InternalServerError("failed to fetch data"))
	}

	data := []entity.ProductResponse{}
	for i := 0; i < len(products); i++ {
		formatRes := entity.FormatProductResponse(products[i])
		data = append(data, formatRes)
	}

	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", data))
}

func (h *productHandler) GetUserProductsController(c echo.Context) error {
	userId := c.Get("currentUser").(entity.User)
	currentUser := userId.Id

	products, err := h.productService.GetAllUserProductsService(currentUser)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to fetch data"))
	}
	formatRes := entity.FormatProductResponse(product)
	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", formatRes))
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
	// fmt.Println("product update = ", productUpdate)

	userId := c.Get("currentUser").(entity.User)
	productUpdate.User = userId

	productUp, err := h.productService.UpdateProductService(idParam, productUpdate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to update data"))
	}

	formatRes := entity.FormatProductResponse(productUp)
	return c.JSON(http.StatusOK, helper.SuccessResponses("success update data", formatRes))
}

func (h *productHandler) DeleteProductController(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed convert id"))
	}

	userId := c.Get("currentUser").(entity.User)
	currentUser := userId.Id

	_, err1 := h.productService.DeleteProductService(idParam, currentUser)

	if err1 != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed delete data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithoutDataResponses("Success delete data"))
}

func (h *productHandler) CreateProductController(c echo.Context) error {
	userInfo := c.Get("currentUser")
	fmt.Println(userInfo.(entity.User).Id)
	userId := userInfo.(entity.User).Id
	newProduct := entity.CreateProduct{}
	if err := c.Bind(&newProduct); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}
	_, err := h.productService.CreateProductService(userId, newProduct)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to insert data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithoutDataResponses("success insert data"))
}
