package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

type AssistHandler struct {
	svc   port.AssistService
	token port.TokenProvider
}

func NewAssistHandler(svc port.AssistService, token port.TokenProvider) *AssistHandler {
	return &AssistHandler{
		svc,
		token,
	}
}

type linkRequest struct {
	LinkID uint64 `param:"id"`
}

func (ah *AssistHandler) RequestAssistedUserLink(c echo.Context) error {
	var req linkRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	auth := ah.token.GetAuth(c)

	err = ah.svc.CreateRequest(auth.ID, req.LinkID, auth.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (ah *AssistHandler) RequestAssistantLink(c echo.Context) error {
	var req linkRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	auth := ah.token.GetAuth(c)

	err = ah.svc.CreateRequest(req.LinkID, auth.ID, auth.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (ah *AssistHandler) AcceptAssistedUserRequest(c echo.Context) error {
	var req linkRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	auth := ah.token.GetAuth(c)

	err = ah.svc.AcceptRequest(auth.ID, req.LinkID, auth.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (ah *AssistHandler) AcceptAssistantRequest(c echo.Context) error {
	var req linkRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	auth := ah.token.GetAuth(c)

	err = ah.svc.AcceptRequest(req.LinkID, auth.ID, auth.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (ah *AssistHandler) DeleteAssistedLink(c echo.Context) error {
	var req linkRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	auth := ah.token.GetAuth(c)

	err = ah.svc.Delete(auth.ID, req.LinkID)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func (ah *AssistHandler) DeleteAssistantLink(c echo.Context) error {
	var req linkRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	auth := ah.token.GetAuth(c)

	err = ah.svc.Delete(req.LinkID, auth.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

type listAssistRequest struct {
	Skip  uint64 `query:"skip"`
	Limit uint64 `query:"limit"`
}

func (ah *AssistHandler) ListAssistants(c echo.Context) error {
	var req listAssistRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	auth := ah.token.GetAuth(c)

	assistants, err := ah.svc.ListAssistants(auth.ID, req.Skip, req.Limit)
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

	auth := ah.token.GetAuth(c)

	assistedUsers, err := ah.svc.ListAssistedUsers(auth.ID, req.Skip, req.Limit)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"assistedUsers": assistedUsers,
	})
}

func (ah *AssistHandler) ListAssistantsRequests(c echo.Context) error {
	var req listAssistRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	auth := ah.token.GetAuth(c)

	assistants, err := ah.svc.ListAssistantsRequests(auth.ID, req.Skip, req.Limit)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"assistants": assistants,
	})
}

func (ah *AssistHandler) ListAssistedUsersRequests(c echo.Context) error {
	var req listAssistRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	auth := ah.token.GetAuth(c)

	assistedUsers, err := ah.svc.ListAssistedUsersRequests(auth.ID, req.Skip, req.Limit)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"assistedUsers": assistedUsers,
	})
}
