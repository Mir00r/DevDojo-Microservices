package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/Mir00r/user-service/configs"
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

// VerifyJWT verifies a JWT and returns the claims if valid
func VerifyJWT(tokenString string) (*JWTClaims, error) {
	//configs.LoadConfig()

	// Debug log to verify JWTSecret
	log.Printf("Received Token: %s", tokenString)
	log.Printf("JWTSecret in VerifyJWT: %s", configs.AppConfig.JWT.Secret)
	secret := []byte(configs.AppConfig.JWT.Secret)
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
func AddClaimsToContext(ctx context.Context, claims *JWTClaims) context.Context {
	return context.WithValue(ctx, "claims", claims)
}

// ExtractClaimsFromContext retrieves JWT claims from the request context
//func ExtractClaimsFromContext(ctx context.Context) (jwt.MapClaims, bool) {
//	claims, ok := ctx.Value("claims").(jwt.MapClaims)
//	return claims, ok
//}

// ExtractClaimsFromContext retrieves JWT claims from the context
func ExtractClaimsFromContext(ctx context.Context) (*JWTClaims, error) {
	claims, ok := ctx.Value("claims").(*JWTClaims) // Type assertion to *JWTClaims
	if !ok {
		return nil, errors.New("invalid or missing JWT claims in context")
	}
	return claims, nil
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

// GenerateRefreshToken generates a secure random token for refresh purposes
func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32) // 32 bytes = 256-bit token
	_, err := rand.Read(bytes)
	if err != nil {
		return "", errors.New("failed to generate refresh token")
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
