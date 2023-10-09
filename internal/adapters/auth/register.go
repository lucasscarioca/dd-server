package auth

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
)

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (uh *UserHandler) register(c echo.Context) error {
	req := new(registerRequest)

	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		return err
	}

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	_, err := uh.svc.Register(&user)
	if err != nil {
	}
	// Hash password

	// Create User if unique

	// Send verification email

	// Create token
	t, err := newToken(req.Email)
	if err != nil {
		return err
	}

	// Create refresh token

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
