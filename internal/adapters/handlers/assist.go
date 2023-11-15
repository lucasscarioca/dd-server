package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

type AssistHandler struct {
	svc port.AssistService
}

func NewAssistHandler(svc port.AssistService) *AssistHandler {
	return &AssistHandler{
		svc,
	}
}

type createAssistRequest struct {
	ID         uint64 `param:"id"`
	AssistedID uint64 `json:"assistedId"`
}

func (ah *AssistHandler) Create(c echo.Context) error {
	var req createAssistRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	err = ah.svc.Create(req.ID, req.AssistedID)
	if err != nil {
		handleError(c, err)
	}

	return c.NoContent(http.StatusCreated)
}

type listAssistRequest struct {
	ID    uint64 `param:"id"`
	Skip  uint64 `query:"skip"`
	Limit uint64 `query:"limit"`
}

func (ah *AssistHandler) ListAssistants(c echo.Context) error {
	var req listAssistRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	assistants, err := ah.svc.ListAssistants(req.ID, req.Skip, req.Limit)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"assistants": assistants,
	})
}

func (ah *AssistHandler) ListAssistedUsers(c echo.Context) error {
	var req listAssistRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	assistedUsers, err := ah.svc.ListAssistedUsers(req.ID, req.Skip, req.Limit)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"assistedUsers": assistedUsers,
	})
}
