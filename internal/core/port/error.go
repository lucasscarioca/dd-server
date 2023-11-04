package port

import (
	"errors"
	"strings"
)

var (
	ErrConflictingData    = errors.New("data conflicts with existing data in unique column")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrNoUpdatedData      = errors.New("no data to update")
	ErrDataNotFound       = errors.New("data not found")
	ErrInvalidToken       = errors.New("access token is invalid")
	ErrExpiredToken       = errors.New("access token has expired")
)

func IsUniqueConstraintViolationError(err error) bool {
	return strings.Contains(err.Error(), "23505")
}
