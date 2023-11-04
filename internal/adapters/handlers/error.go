package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

var errorStatusMap = map[error]int{
	port.ErrConflictingData:    http.StatusConflict,
	port.ErrInvalidCredentials: http.StatusBadRequest,
}

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func newErrorResponse(message string) errorResponse {
	return errorResponse{
		Success: false,
		Message: message,
	}
}

func handleError(c echo.Context, err error) error {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errResp := newErrorResponse(err.Error())
	return c.JSON(statusCode, errResp)
}

func validationError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
}
