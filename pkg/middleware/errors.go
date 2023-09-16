package middleware

import (
	"errors"
)

var (
	// ErrorInvalidToken - is a constant that represents an invalid token error.
	ErrorInvalidToken = errors.New("authorization failed: invalid token")

	// ErrorExpiredToken - is a constant that represents an expired token error.
	ErrorExpiredToken = errors.New("authorization failed: expired token")

	// ErrorInvalidCredentials - is a constant that represents an invalid credentials error.
	ErrorInvalidCredentials = errors.New("authorization failed: invalid credentials")

	// ErrorUnauthorizedUser - is a constant that represents an unauthorized user error.
	ErrorUnauthorizedUser = errors.New("authorization failed: unauthorized user")
)
