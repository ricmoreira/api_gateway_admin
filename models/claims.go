package models

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	// ErrorResponse represents the token model for token response messages
	Claims struct {
		Username string `json:"username"`
		Roles []Role `json:"roles"`

		jwt.StandardClaims
	}
)
