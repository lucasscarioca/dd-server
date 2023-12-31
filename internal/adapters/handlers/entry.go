package handlers

import (
	"encoding/json"
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
	Title     string                 `json:"title"`
	Content   string                 `json:"content"`
	Status    string                 `json:"status"`
	Configs   map[string]interface{} `json:"configs"` //TODO: type configs according to app
	CreatedAt string                 `json:"createdAt"`
}

func (eh *EntryHandler) Create(c echo.Context) error {
	var req createEntryRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	parsedConfigs, err := json.Marshal(req.Configs)
	if err != nil {
		return validationError(c, err)
	}

	var createdAt time.Time
	if len(req.CreatedAt) > 0 {
		createdAt, err = time.Parse(time.DateOnly, req.CreatedAt)
		if err != nil {
			createdAt = time.Now()
		}
	} else {
		createdAt = time.Now()
	}

	auth := eh.token.GetAuth(c)

	newEntry := domain.Entry{
		Title:     req.Title,
		Content:   req.Content,
		UserID:    auth.ID,
		Status:    req.Status,
		Configs:   parsedConfigs,
		CreatedAt: createdAt,
		UpdatedAt: time.Now(),
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

	auth := eh.token.GetAuth(c)

	entries, err := eh.svc.List(auth.ID, req.Skip, req.Limit, req.Date)
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

	auth := eh.token.GetAuth(c)

	entry, err := eh.svc.Find(auth.ID, req.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"entry": entry,
	})
}

type updateEntryRequest struct {
	ID      uint64                 `param:"id"`
	Title   string                 `json:"title,omitempty"`
	Content string                 `json:"content,omitempty"`
	Status  string                 `json:"status,omitempty"`
	Configs map[string]interface{} `json:"configs,omitempty"`
}

func (eh *EntryHandler) Update(c echo.Context) error {
	var req updateEntryRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	parsedConfigs, err := json.Marshal(req.Configs)
	if err != nil {
		return validationError(c, err)
	}

	auth := eh.token.GetAuth(c)

	newEntry := domain.Entry{
		ID:        req.ID,
		Title:     req.Title,
		Content:   req.Content,
		UserID:    auth.ID,
		Status:    req.Status,
		Configs:   parsedConfigs,
		UpdatedAt: time.Now(),
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

	auth := eh.token.GetAuth(c)

	err = eh.svc.Delete(auth.ID, req.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}
