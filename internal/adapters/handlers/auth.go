package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

type AuthHandler struct {
	svc port.AuthService
}

func NewAuthHandler(svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc,
	}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ah *AuthHandler) Login(c echo.Context) error {
	req := new(loginRequest)

	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		return err
	}

	t, err := ah.svc.Login(req.Email, req.Password)
	if err != nil {
		return handleError(c, err)
	}

	// Create refresh token

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ah *AuthHandler) Register(c echo.Context) error {
	req := new(registerRequest)
	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		return err
	}

	t, err := ah.svc.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return handleError(c, err)
	}

	// Create refresh token

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

type forgotRequest struct {
	Email string `json:"email"`
}

func (ah *AuthHandler) Forgot(c echo.Context) error {
	req := new(forgotRequest)

	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		return err
	}

	err := ah.svc.Forgot(req.Email)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

type resetRequest struct {
	Password string `json:"password"`
}

func (ah *AuthHandler) Reset(c echo.Context) error {
	token := c.Param("token")
	req := new(resetRequest)

	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		return err
	}

	err := ah.svc.Reset(req.Password, token)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}
