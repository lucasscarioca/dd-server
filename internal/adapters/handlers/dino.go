package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

type DinoHandler struct {
	svc   port.DinoService
	token port.TokenProvider
}

func NewDinoHandler(svc port.DinoService, token port.TokenProvider) *DinoHandler {
	return &DinoHandler{
		svc,
		token,
	}
}

type createDinoRequest struct {
	Name    string                 `json:"name"`
	Avatar  string                 `json:"avatar"`
	Configs map[string]interface{} `json:"configs"` //TODO: type configs according to app
}

func (dh *DinoHandler) Create(c echo.Context) error {
	var req createDinoRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	parsedConfigs, err := json.Marshal(req.Configs)
	if err != nil {
		return validationError(c, err)
	}

	auth := dh.token.GetAuth(c)

	newDino := domain.Dino{
		Name:      req.Name,
		Avatar:    req.Avatar,
		Configs:   parsedConfigs,
		UserID:    auth.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	dino, err := dh.svc.Create(&newDino)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"dino": dino,
	})
}

type listDinoRequest struct {
	Skip  uint64 `query:"skip"`
	Limit uint64 `query:"limit"`
}

func (dh *DinoHandler) List(c echo.Context) error {
	var req listDinoRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	auth := dh.token.GetAuth(c)

	dinos, err := dh.svc.List(auth.ID, req.Skip, req.Limit)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"dinos": dinos,
	})
}

type findDinoRequest struct {
	ID uint64 `param:"id"`
}

func (dh *DinoHandler) Find(c echo.Context) error {
	var req findDinoRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	auth := dh.token.GetAuth(c)

	dino, err := dh.svc.Find(auth.ID, req.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"dino": dino,
	})
}

type updateDinoRequest struct {
	ID      uint64                 `param:"id"`
	Name    string                 `json:"name,omitempty"`
	Avatar  string                 `json:"avatar,omitempty"`
	Configs map[string]interface{} `json:"configs,omitempty"`
}

func (dh *DinoHandler) Update(c echo.Context) error {
	var req updateDinoRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	parsedConfigs, err := json.Marshal(req.Configs)
	if err != nil {
		return validationError(c, err)
	}

	auth := dh.token.GetAuth(c)

	newDino := domain.Dino{
		ID:        req.ID,
		Name:      req.Name,
		Avatar:    req.Avatar,
		Configs:   parsedConfigs,
		UserID:    auth.ID,
		UpdatedAt: time.Now(),
	}

	dino, err := dh.svc.Update(&newDino)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"dino": dino,
	})
}

type deleteDinoRequest struct {
	ID uint64 `param:"id"`
}

func (dh *DinoHandler) Delete(c echo.Context) error {
	var req deleteDinoRequest
	err := c.Bind(&req)
	if err != nil {
		return validationError(c, err)
	}

	auth := dh.token.GetAuth(c)

	err = dh.svc.Delete(auth.ID, req.ID)
	if err != nil {
		return handleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}
