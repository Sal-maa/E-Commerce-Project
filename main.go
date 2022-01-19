package main

import (
	"os"

	"github.com/Sal-maa/E-Commerce-Project/helper"
	"github.com/labstack/echo/v4"
)

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := helper.InitDB(connectionString)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	router.UserRouter(e, db, jwtSecret)

	e.Logger.Fatal(e.Start(":8080"))

}
