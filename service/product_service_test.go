package service_test

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Sal-maa/E-Commerce-Project/entity"
	"github.com/Sal-maa/E-Commerce-Project/helper"
	"github.com/Sal-maa/E-Commerce-Project/repo"
	"github.com/Sal-maa/E-Commerce-Project/router"
	"github.com/Sal-maa/E-Commerce-Project/service"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func InitTestEchoAPIProduct() (*echo.Echo, *sql.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("DB_CONNECTION_TEST")
	db, err := helper.InitDB(connectionString)
	if err != nil {
		panic(err)
	}
	e := echo.New()
	router.UserRouter(db)
	return e, db
}

func TestCreateProduct(t *testing.T) {
	e, db := InitTestEchoAPIProduct()
	defer db.Close()

	productRepository := repo.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	t.Run("Test Create Success", func(t *testing.T) {

		input := entity.CreateProduct{
			CategoryId: 3,
			Name:       "buku apik",
			Deskripsi:  "buku bacaan",
			Gambar:     "www.gambar.com",
			Harga:      9000,
			Stock:      21,
		}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")

		result, err := productService.CreateProductService(1, input)
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
	productRepository := repo.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	t.Run("TestGetAllProducts", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		result, err := productService.GetAllUserProductsService(4)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetAllProducts", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		result, err := productService.GetAllProductsService()
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

		result, err := productService.GetProductByIdService(1000)
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
	productRepository := repo.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	t.Run("TestUpdateSuccess", func(t *testing.T) {
		input := entity.EditProduct{
			CategoryId: 3,
			Name:       "buku apik",
			Deskripsi:  "buku bacaan",
			Gambar:     "www.gambar.com",
			Harga:      19000,
			Stock:      21,
		}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")

		result, err := productService.UpdateProductService(5, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestUpdateError", func(t *testing.T) {
		input := entity.EditProduct{
			CategoryId: 3,
			Name:       "buku apik",
			Deskripsi:  "buku bacaan",
			Gambar:     "www.gambar.com",
			Harga:      9000,
			Stock:      21,
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

	productRepository := repo.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	t.Run("TestDeleteSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")

		result, err := productService.DeleteProductService(1, 2)
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

		result, err := productService.DeleteProductService(99, 100)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}
