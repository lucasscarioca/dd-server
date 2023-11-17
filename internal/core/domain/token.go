package domain

import "github.com/golang-jwt/jwt/v5"

type Token struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewTokenClaims(id uint64, email string, exp jwt.NumericDate) jwt.Claims {
	return Token{
		id,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: &exp,
		},
	}
}
