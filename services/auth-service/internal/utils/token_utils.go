package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var ResetTokenSecret = []byte("7CD0WF6Yuu") // Replace with a secure key

// VerifyPasswordResetToken verifies the password reset token and extracts the user ID
func VerifyPasswordResetToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return ResetTokenSecret, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	// Ensure the token is for password reset
	if purpose, ok := claims["purpose"].(string); !ok || purpose != "password_reset" {
		return "", errors.New("invalid token purpose")
	}

	// Check token expiration
	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return "", errors.New("token expired")
	}

	// Extract and return the user ID
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	return userID, nil
}

// GeneratePasswordResetToken generates a JWT for password reset
func GeneratePasswordResetToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"purpose": "password_reset",                      // Purpose for password reset
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//log.Printf("Env reset secret key: %v\n", config.GetConfig().ResetTokenSecret)
	return token.SignedString(ResetTokenSecret)
}
