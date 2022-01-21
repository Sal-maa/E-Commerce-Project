package main

import (
	"log"
	"os"

	"github.com/Sal-maa/E-Commerce-Project/helper"
	"github.com/Sal-maa/E-Commerce-Project/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("DB_CONNECTION")
	db, err := helper.InitDB(connectionString)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	router.UserRouter(e, db)

	e.Logger.Fatal(e.Start(":8080"))

}
