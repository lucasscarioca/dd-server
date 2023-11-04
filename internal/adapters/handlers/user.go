package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
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

type findRequest struct {
	ID uint64 `param:"id"`
}

func (uh *UserHandler) Find(c echo.Context) error {
	var req findRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	user, err := uh.svc.Find(req.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

type updateRequest struct {
	ID      uint64 `param:"id"`
	Name    string `json:"name,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
	Email   string `json:"email,omitempty"`
	Configs any    `json:"configs,omitempty"`
}

func (uh *UserHandler) Update(c echo.Context) error {
	var req updateRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	newUser := domain.User{
		ID:      req.ID,
		Name:    req.Name,
		Avatar:  req.Avatar,
		Email:   req.Email,
		Configs: req.Configs,
	}

	user, err := uh.svc.Update(&newUser)
	if err != nil {
		handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

type deleteRequest struct {
	ID uint64 `param:"id"`
}

func (uh *UserHandler) Delete(c echo.Context) error {
	var req deleteRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = uh.svc.Delete(req.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}
