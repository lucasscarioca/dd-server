package auth

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type forgotRequest struct {
	Email string `json:"email"`
}

func forgot(c echo.Context) error {
	req := new(forgotRequest)

	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		return err
	}

	// Find User by email

	// Create token

	// Update User with token

	// Send email

	return c.NoContent(http.StatusOK)
}
