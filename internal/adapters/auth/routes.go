package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

func Routes(e *echo.Echo) {
	e.POST("/forgot", forgot)
	e.PUT("/reset/:token", reset)
	e.GET("/verify/:token", verify)
	e.POST("/refresh", refresh)
}

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}
