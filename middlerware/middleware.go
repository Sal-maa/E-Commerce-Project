package middlerware

import (
	"net/http"
	"strings"

	"github.com/Sal-maa/E-Commerce-Project/helper"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(authService JWTService, userService service.UserService, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.UnauthorizedResponses("header not")
			return c.JSON(http.StatusUnauthorized, response)
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.UnauthorizedResponses("token Unauthorized")
			return c.JSON(http.StatusUnauthorized, response)
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response := helper.UnauthorizedResponses("payload Unauthorized")
			return c.JSON(http.StatusUnauthorized, response)
		}

		userName := string(payload["userName"].(string))
		user, err := userService.GetUserByNameService(userName)
		if err != nil {
			response := helper.UnauthorizedResponses("user name Unauthorized")
			return c.JSON(http.StatusUnauthorized, response)
		}
		c.Set("currentUser", user)
		return next(c)
	}
}
