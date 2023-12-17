package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	UserID string
	jwt.RegisteredClaims
}
