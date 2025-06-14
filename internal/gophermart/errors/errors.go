package errors

import "errors"

var (
	ErrInitDB         = errors.New("failed to init database")
	ErrEmptySecretKey = errors.New("empty secret key for JWT")
	ErrHTTPServer     = errors.New("failed to start HTTP server")
)
