package utils

import (
	"context"
	"errors"
	config "github.com/Mir00r/auth-service/internal/configs"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

// JWTClaims defines the claims used in the JWT.
type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token with the provided user ID and email.
func GenerateJWT(userID, email, secret string, expiry time.Duration) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// VerifyJWT validates a JWT token and returns the claims if valid.
//func VerifyJWT(tokenString, secret string) (*JWTClaims, error) {
//	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, errors.New("unexpected signing method")
//		}
//		return []byte(secret), nil
//	})
//
//	if err != nil {
//		return nil, err
//	}
//
//	claims, ok := token.Claims.(*JWTClaims)
//	if !ok || !token.Valid {
//		return nil, errors.New("invalid token")
//	}
//
//	return claims, nil
//}

// VerifyJWT verifies a JWT and returns the claims if valid
func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		config.LoadConfig()

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return config.GetConfig().JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// AddClaimsToContext adds JWT claims to the request context
func AddClaimsToContext(ctx context.Context, claims jwt.MapClaims) context.Context {
	return context.WithValue(ctx, "claims", claims)
}

// ExtractClaimsFromContext retrieves JWT claims from the request context
func ExtractClaimsFromContext(ctx context.Context) (jwt.MapClaims, bool) {
	claims, ok := ctx.Value("claims").(jwt.MapClaims)
	return claims, ok
}

// ValidateBasicAuth validates the username and password for Basic Authentication
func ValidateBasicAuth(username, password string) bool {
	// Fetch the expected username and password from environment variables
	expectedUsername := os.Getenv("BASIC_AUTH_USERNAME")
	expectedPassword := os.Getenv("BASIC_AUTH_PASSWORD")

	// Check if the provided username and password match the expected values
	if username == expectedUsername && password == expectedPassword {
		return true
	}

	return false
}
