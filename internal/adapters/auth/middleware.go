package auth

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Middleware() echo.MiddlewareFunc {
	// Configure middleware with the custom claims type
	middlewareConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(claims)
		},
		SigningKey: []byte("secret"),
	}

	return echojwt.WithConfig(middlewareConfig)
}
