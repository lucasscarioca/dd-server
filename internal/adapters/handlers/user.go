package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

type listRequest struct {
	Skip  uint64 `query:"skip"`
	Limit uint64 `query:"limit"`
}

func (uh *UserHandler) List(c echo.Context) error {
	var options listRequest
	err := c.Bind(&options)
	if err != nil {
		return err
	}

	users, err := uh.svc.List(options.Skip, options.Limit)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"users": users,
	})
}
