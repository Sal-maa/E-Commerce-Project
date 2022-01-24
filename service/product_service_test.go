package service_test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/helper"
	"github.com/Sal-maa/E-Commerce-Project/router"
	"github.com/Sal-maa/E-Commerce-Project/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func InitTestEchoAPIProduct() (*echo.Echo, *sql.DB) {
	jwtSecret := os.Getenv("JWT_SECRET")
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := helper.InitDB(connectionString)
	if err != nil {
		panic(err)
	}
	e := echo.New()
	router.UserRouter(e, db, jwtSecret)
	return e, db
}

func TestCreateProduct(t *testing.T) {
	e, db := InitTestEchoAPIProduct()
	defer db.Close()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	t.Run("Test Create Success", func(t *testing.T) {

		input := entity.CreateProductRequest{
			UserId: 3,
			Name:   "biscuit",
			Price:  3000,
			Stock:  19,
		}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")

		result, err := productService.CreateProductService(input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestGetProduct(t *testing.T) {
	// setting controller
	e, db := InitTestEchoAPIProduct()
	defer db.Close()
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	t.Run("TestGetAll", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		result, err := productService.GetProductsService()
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyIdSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")

		result, err := productService.GetProductByIdService(7)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyIdError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")

		result, err := productService.GetProductByIdService(1)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestUpdateProduct(t *testing.T) {
	// setting controller
	e, db := InitTestEchoAPIProduct()
	defer db.Close()
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	t.Run("TestUpdateSuccess", func(t *testing.T) {
		input := entity.EditProductRequest{
			UserId: 3,
			Name:   "kopi",
			Price:  10000,
			Stock:  8,
		}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")

		result, err := productService.UpdateProductService(8, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestUpdateError", func(t *testing.T) {
		input := entity.EditProductRequest{
			UserId: 3,
			Name:   "soap",
			Price:  6000,
			Stock:  22,
		}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")

		result, err := productService.UpdateProductService(6, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestDeleteProduct(t *testing.T) {
	// setting controller
	e, db := InitTestEchoAPIProduct()
	defer db.Close()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	t.Run("TestDeleteSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")

		result, err := productService.DeleteProductService(10)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestDeleteError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")

		result, err := productService.DeleteProductService(5)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}
