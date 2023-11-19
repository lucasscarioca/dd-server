package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

type UserHandler struct {
	svc   port.UserService
	token port.TokenProvider
}

func NewUserHandler(svc port.UserService, token port.TokenProvider) *UserHandler {
	return &UserHandler{
		svc,
		token,
	}
}

type listUserRequest struct {
	Skip  uint64 `query:"skip"`
	Limit uint64 `query:"limit"`
}

func (uh *UserHandler) List(c echo.Context) error {
	var options listUserRequest
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

func (uh *UserHandler) Profile(c echo.Context) error {
	auth := uh.token.GetAuth(c)

	user, err := uh.svc.Find(auth.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

type updateRequest struct {
	Name    string                 `json:"name,omitempty"`
	Avatar  string                 `json:"avatar,omitempty"`
	Email   string                 `json:"email,omitempty"`
	Configs map[string]interface{} `json:"configs,omitempty"`
}

func (uh *UserHandler) Update(c echo.Context) error {
	var req updateRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	parsedConfigs, err := json.Marshal(req.Configs)
	if err != nil {
		return validationError(c, err)
	}

	auth := uh.token.GetAuth(c)

	newUser := domain.User{
		ID:      auth.ID,
		Name:    req.Name,
		Avatar:  req.Avatar,
		Email:   req.Email,
		Configs: parsedConfigs,
	}

	user, err := uh.svc.Update(&newUser)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

func (uh *UserHandler) Delete(c echo.Context) error {
	auth := uh.token.GetAuth(c)

	err := uh.svc.Delete(auth.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}
