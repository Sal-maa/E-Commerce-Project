package main

import (
	"log"
	"os"

	"github.com/Sal-maa/E-Commerce-Project/helper"
	"github.com/Sal-maa/E-Commerce-Project/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := helper.InitDB(connectionString)
	if err != nil {
		panic(err)
	}

	// e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e := router.UserRouter(db)

	e.Logger.Fatal(e.Start(":8080"))

}
