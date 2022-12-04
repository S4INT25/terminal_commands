package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"terminal_commands/models"
	"time"
)

const secretKey = "TI3CYPlE3PnddyTXYvEK3jKxVcew99En"

func GenerateJwtToken(user models.User) string {

	secretKey := []byte(secretKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Date(2024, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		panic(fmt.Sprintf("Failed to sign key %v", err))
	}

	return tokenString

}

func ValidateToken(jwtToken string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}

}
