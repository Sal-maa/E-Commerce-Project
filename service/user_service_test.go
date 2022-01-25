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
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func InitEchoTestAPIUser() (*echo.Echo, *sql.DB) {
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

func TestCreate(t *testing.T) {
	e, db := InitEchoTestAPIUser()
	defer db.Close()

	userRepository := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	t.Run("Check Username", func(t *testing.T) {
		input := entity.CreateUserRequest{
			Username: "suzana",
			Email:    "suz@na",
			Password: "suzana",
			Address:  "dirumah",
			Phone:    "90989787",
		}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")

		result, err := userService.CheckUserName(input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)

	})
	t.Run("Test Create Success", func(t *testing.T) {
		input := entity.CreateUserRequest{
			Username: "suzana",
			Email:    "suz@na",
			Password: "suzana",
			Address:  "dirumah",
			Phone:    "90989787",
		}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")

		result, err := userService.CreateUserService(input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestLogin(t *testing.T) {
	e, db := InitEchoTestAPIUser()
	defer db.Close()

	userRepository := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	t.Run("Test Login Success", func(t *testing.T) {

		input := entity.LoginUserRequest{
			Username: "suzana",
			Password: "suzana",
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		result, err := userService.LoginUserService(input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("Test Login Error", func(t *testing.T) {
		input := entity.LoginUserRequest{
			Username: "suzana",
			Password: "543",
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		result, err := userService.LoginUserService(input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("Save Token", func(t *testing.T) {

		Token := "string"
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		result, err := userService.SaveTokenService(Token)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestGet(t *testing.T) {
	// setting controller
	e, db := InitEchoTestAPIUser()
	defer db.Close()
	userRepository := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	t.Run("TestGetbyIdSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		result, err := userService.GetUserByIdService(23)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyIdError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		result, err := userService.GetUserByIdService(20)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyNameSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		result, err := userService.GetUserByNameService("suzana")
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})

	t.Run("TestGetbyNameError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		result, err := userService.GetUserByNameService("aku")
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestDelete(t *testing.T) {
	// setting controller
	e, db := InitEchoTestAPIUser()

	userRepository := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	t.Run("TestDeleteSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		result, err := userService.DeleteUserService(4)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestDeleteError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		result, err := userService.DeleteUserService(20)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestUpdate(t *testing.T) {
	// setting controller
	e, db := InitEchoTestAPIUser()

	userRepository := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	t.Run("TestUpdateSuccess", func(t *testing.T) {
		input := entity.EditUserRequest{
			Username: "suzana",
			Email:    "suz@ni",
			Password: "123susi",
			Address:  "dirumah aja",
		}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		result, err := userService.UpdateUserService(3, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestUpdateError", func(t *testing.T) {
		input := entity.EditUserRequest{
			Username: "suzene",
			Email:    "suz@ne",
			Password: "123suse",
			Address:  "dirumah ye",
		}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")

		result, err := userService.UpdateUserService(20, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}
