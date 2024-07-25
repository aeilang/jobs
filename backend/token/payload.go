package token

import (
	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}