package port

import (
	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
)

type TokenProvider interface {
	Create(id uint64, email string) (string, error)
	GetAuth(c echo.Context) *domain.Token
}

type AuthService interface {
	Login(email, password string) (string, error)
	Register(name, email, password string) (string, error)
	Forgot(email string) error
	Reset(password, token string) error
}
