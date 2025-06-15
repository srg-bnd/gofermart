package services

import (
	"errors"
)

var (
	ErrInvalidCreds = errors.New("invalid credentials")
	ErrUserExists   = errors.New("user already exists")
	ErrJWTToken     = errors.New("invalid JWT token")
)
