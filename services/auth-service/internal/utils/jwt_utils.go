package utils

import (
	"context"
	"errors"
	config "github.com/Mir00r/auth-service/internal/configs"
	"github.com/golang-jwt/jwt/v4"
	"log"
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
	log.Printf("JWTSecret While Generating Token: %v", secret)
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
	log.Printf("Generated Token: %s", signedToken)
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
func VerifyJWT(tokenString string) (*JWTClaims, error) {
	config.LoadConfig()

	// Debug log to verify JWTSecret
	log.Printf("Received Token: %s", tokenString)
	log.Printf("JWTSecret in VerifyJWT: %s", config.GetConfig().JWTSecret)
	secret := []byte(config.GetConfig().JWTSecret)
	log.Printf("Using JWT Secret After Conevrtion Byte Data Type: %s", secret)

	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		// Return the secret for verification
		return secret, nil
	})

	if err != nil {
		log.Printf("JWT Parsing Error: %v", err)
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
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
