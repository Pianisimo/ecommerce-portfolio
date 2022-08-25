package models

import (
	"ecommerce/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	RefreshedTokenValidTime = time.Hour * 72
	AuthTokenValidTime      = time.Minute * 15
)

type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	Csrf string `json:"csrf"`
}

// GenerateCSRFSecret Generates a random string to be used as a csrf secret
func GenerateCSRFSecret() (string, error) {
	return utils.GenerateRandomString(32)
}
