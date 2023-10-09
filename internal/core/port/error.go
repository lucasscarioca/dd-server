package port

import (
	"errors"
	"strings"
)

var (
	ErrConflictingData    = errors.New("data conflicts with existing data in unique column")
	ErrInvalidCredentials = errors.New("invalid email or password")
)

func IsUniqueConstraintViolationError(err error) bool {
	return strings.Contains(err.Error(), "23505")
}
