package middleware

import (
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID          string
	Fullname    string
	Username    string
	Email       string
	DateOfBirth string
	Phone       string
	Roles       []string
}

type SignedParams struct {
	User *User
	jwt.RegisteredClaims
	Roles []string
}

// Store Structs
type CtxKey interface{}
type CtxValues struct {
	m map[string]interface{}
}
