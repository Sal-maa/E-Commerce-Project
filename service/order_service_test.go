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

func InitTestEchoAPIOrder() (*echo.Echo, *sql.DB) {
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

func TestCreateOrder(t *testing.T) {
	e, db := InitTestEchoAPIOrder()
	defer db.Close()

	orderRepository := repo.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)

	t.Run("Test Create Success", func(t *testing.T) {

		input := entity.CreateOrderRequest{
			CartId: []int{1, 2, 4},
			Address: entity.Address{
				Street: "jalan",
				City:   "kota",
				State:  "negara",
				Zip:    1234,
			},
			CreditCard: entity.CreditCard{
				Type:   "tipe",
				Name:   "nama",
				Number: "1234-1234-1234-1234",
				CVV:    1234,
			},
			Total: 1234566,
		}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/orders")

		result, err := orderService.CreateOrderService(input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestGetOrder(t *testing.T) {
	// setting controller
	e, db := InitTestEchoAPIOrder()
	defer db.Close()
	orderRepository := repo.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)

	t.Run("TestGetAll", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/orders")

		result, err := orderService.GetOrderService(1)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyIdSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/orders/:id")
		context.SetParamNames("id")

		result, err := orderService.GetOrderByIdService(3, 1)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyIdError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/orders/:id")
		context.SetParamNames("id")

		result, err := orderService.GetOrderByIdService(3, 1)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestUpdateOrder(t *testing.T) {
	// setting controller
	e, db := InitTestEchoAPIOrder()
	defer db.Close()
	orderRepository := repo.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)

	input := entity.EditOrderRequest{
		StatusOrder: "Cancelled",
	}

	t.Run("TestUpdateSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/orders/:id")
		context.SetParamNames("id")

		result, err := orderService.UpdateOrderService(1, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestUpdateError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/carts/:id")
		context.SetParamNames("id")

		result, err := orderService.UpdateOrderService(0, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}
