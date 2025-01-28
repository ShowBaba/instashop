package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthTokenJwtClaim struct {
	Email string
	ID    uint
	Role  string
	jwt.StandardClaims
}

type TokenStruct struct {
	UserID    uint
	Token     int
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
