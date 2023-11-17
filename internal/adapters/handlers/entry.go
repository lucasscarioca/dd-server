package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

type EntryHandler struct {
	svc   port.EntryService
	token port.TokenProvider
}

func NewEntryHandler(svc port.EntryService, token port.TokenProvider) *EntryHandler {
	return &EntryHandler{
		svc,
		token,
	}
}

type createEntryRequest struct {
	Title   string         `json:"title"`
	Content string         `json:"content"`
	Status  string         `json:"status"`
	Configs map[string]any `json:"configs"` //TODO: type configs according to app
}

func (eh *EntryHandler) Create(c echo.Context) error {
	var req createEntryRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	newEntry := domain.Entry{
		Title:   req.Title,
		Content: req.Content,
		Status:  req.Status,
		Configs: req.Configs,
	}
	entry, err := eh.svc.Create(&newEntry)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"entry": entry,
	})
}

type listEntryRequest struct {
	Skip  uint64 `query:"skip"`
	Limit uint64 `query:"limit"`
	Date  string `query:"date"`
}

func (eh *EntryHandler) List(c echo.Context) error {
	var req listEntryRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	parsedDate, err := time.Parse(time.DateOnly, req.Date)
	if err != nil {
		return validationError(c, err)
	}

	auth := eh.token.GetAuth(c)

	entries, err := eh.svc.List(auth.ID, req.Skip, req.Limit, parsedDate)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"entries": entries,
	})
}

type findEntryRequest struct {
	ID uint64 `param:"id"`
}

func (eh *EntryHandler) Find(c echo.Context) error {
	var req findEntryRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	entry, err := eh.svc.Find(req.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"entry": entry,
	})
}

type updateEntryRequest struct {
	ID      uint64         `param:"id"`
	Title   string         `json:"title,omitempty"`
	Content string         `json:"content,omitempty"`
	Status  string         `json:"status,omitempty"`
	Configs map[string]any `json:"configs,omitempty"`
}

func (eh *EntryHandler) Update(c echo.Context) error {
	var req updateEntryRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	newEntry := domain.Entry{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
		Status:  req.Status,
		Configs: req.Configs,
	}

	entry, err := eh.svc.Update(&newEntry)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"entry": entry,
	})
}

type deleteEntryRequest struct {
	ID uint64 `param:"id"`
}

func (eh *EntryHandler) Delete(c echo.Context) error {
	var req deleteEntryRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	err = eh.svc.Delete(req.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}
