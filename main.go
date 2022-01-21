package main

import (
	"fmt"

	"github.com/Sal-maa/E-Commerce-Project/helper"
	"github.com/Sal-maa/E-Commerce-Project/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	connectionString := fmt.Sprintf("root:123456789@tcp(localhost:3306)/e-commerce?charset=utf8&parseTime=True&loc=Local")
	db, err := helper.InitDB(connectionString)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	router.UserRouter(e, db)

	e.Logger.Fatal(e.Start(":8080"))

}
