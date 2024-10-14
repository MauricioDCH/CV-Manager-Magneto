package jwtService

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = "cv-manager-key"

func GenerateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"sub": strconv.Itoa(userID),
		"exp": time.Now().Add(20 * time.Minute).Unix(), // Expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
