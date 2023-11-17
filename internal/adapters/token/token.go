package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
)

type TokenProvider struct {
	expiration jwt.NumericDate
	key        []byte
}

func NewTokenProvider(hours time.Duration, key string) *TokenProvider {
	return &TokenProvider{
		expiration: *jwt.NewNumericDate(time.Now().Add(time.Hour * hours)),
		key:        []byte(key),
	}
}

func (tp *TokenProvider) Create(id uint64, email string) (string, error) {
	// Set custom claims
	tokenClaims := domain.NewTokenClaims(id, email, tp.expiration)

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	// Generate encoded token and send it as response
	t, err := token.SignedString(tp.key)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (tp *TokenProvider) Authenticate() echo.MiddlewareFunc {
	// Configure middleware with the custom claims type
	middlewareConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(domain.Token)
		},
		SigningKey: tp.key,
	}

	return echojwt.WithConfig(middlewareConfig)
}

func (tp *TokenProvider) GetAuth(c echo.Context) *domain.Token {
	auth := c.Get("user").(*jwt.Token)
	return auth.Claims.(*domain.Token)
}
