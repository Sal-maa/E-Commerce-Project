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

func InitTestEchoAPICart() (*echo.Echo, *sql.DB) {
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

func TestCreateCart(t *testing.T) {
	e, db := InitTestEchoAPICart()
	defer db.Close()

	cartRepository := repo.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository)

	t.Run("Test Create Success", func(t *testing.T) {

		input := entity.CreateCartRequest{
			ProductId: 1,
			Qty:       2,
			Subtotal:  90000,
		}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts")

		result, err := cartService.CreateCartService(input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestGetCart(t *testing.T) {
	// setting controller
	e, db := InitTestEchoAPICart()
	defer db.Close()
	cartRepository := repo.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository)

	user := entity.User{
		Username: "sintia",
		Email:    "halo",
		Password: "12mks",
		Address:  "disana",
		Phone:    "kskxnsa",
	}

	t.Run("TestGetAll", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/carts")

		result, err := cartService.GetAllCartsService(user)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyIdSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/carts/:id")
		context.SetParamNames("id")

		result, err := cartService.GetCartByIdService(1)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyIdError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/carts/:id")
		context.SetParamNames("id")

		result, err := cartService.GetCartByIdService(100)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestUpdateCart(t *testing.T) {
	// setting controller
	e, db := InitTestEchoAPICart()
	defer db.Close()
	cartRepository := repo.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository)

	t.Run("TestUpdateSuccess", func(t *testing.T) {
		input := entity.EditCartRequest{
			Qty:      21,
			Subtotal: 800000,
		}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/carts/:id")
		context.SetParamNames("id")

		result, err := cartService.UpdateCartService(8, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestUpdateError", func(t *testing.T) {
		input := entity.EditCartRequest{
			Qty:      21,
			Subtotal: 800000,
		}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/carts/:id")
		context.SetParamNames("id")

		result, err := cartService.UpdateCartService(600, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestDeleteCart(t *testing.T) {
	// setting controller
	e, db := InitTestEchoAPICart()
	defer db.Close()

	cartRepository := repo.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository)

	t.Run("TestDeleteSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/carts/:id")
		context.SetParamNames("id")

		result, err := cartService.DeleteCartService(1, 2)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestDeleteError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/carts/:id")
		context.SetParamNames("id")

		result, err := cartService.DeleteCartService(99, 100)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}
