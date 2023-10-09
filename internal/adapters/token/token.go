package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
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

type claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (tp *TokenProvider) CreateToken(email string) (string, error) {
	// Set custom claims
	tokenClaims := &claims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: &tp.expiration,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	// Generate encoded token and send it as response
	t, err := token.SignedString(tp.key)
	if err != nil {
		return "", err
	}

	return t, nil
}
