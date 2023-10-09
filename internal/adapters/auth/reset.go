package auth

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type resetRequest struct {
	Password string `json:"password"`
}

func reset(c echo.Context) error {
	// token := c.Param("token")
	req := new(resetRequest)

	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		return err
	}

	// Find User by token

	// Hash new password

	// Update User and clear remember token

	return nil
}
