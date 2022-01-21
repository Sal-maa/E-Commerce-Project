package router

import (
	"database/sql"

	"github.com/Sal-maa/E-Commerce-Project/handler"
	"github.com/Sal-maa/E-Commerce-Project/middleware"
	"github.com/Sal-maa/E-Commerce-Project/repo"
	"github.com/Sal-maa/E-Commerce-Project/service"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func UserRouter(e *echo.Echo, db *sql.DB) {
	e.Pre(echoMiddleware.RemoveTrailingSlash(), echoMiddleware.Logger())
	authService := middleware.AuthService()
	// Route User
	userRepository := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(authService, userService)

	e.POST("/login", userHandler.AuthController)
	e.POST("/users", userHandler.CreateUserController)
	e.GET("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.GetUserController))
	e.PUT("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.UpdateUserController))
	e.DELETE("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.DeleteUserController))

	// Route Cart
	cartRepository := repo.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository)
	cartHandler := handler.NewCartHandler(cartService)

	e.POST("/carts", middleware.AuthMiddleware(authService, userService, cartHandler.CreateCartController))
	// e.GET("/carts", middleware.AuthMiddleware(authService, userService,cartHandler.GetCartsController)
	// e.PUT("/carts/:id", middleware.AuthMiddleware(authService, userService, cartHandler.UpdateCartController))
	// e.DELETE("/carts/:id", middleware.AuthMiddleware(authService, userService, cartHandler.DeleteCartController))


	// Route product
	productRepository := repo.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	e.GET("/products", middleware.AuthMiddleware(authService, userService, productHandler.GetAllProductsController))
	e.GET("/products/:id",  middleware.AuthMiddleware(authService, userService, productHandler.GetProductByIdController))

}
