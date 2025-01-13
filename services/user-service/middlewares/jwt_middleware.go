package middlewares

import (
	"github.com/Mir00r/user-service/config"
	"github.com/Mir00r/user-service/constants"
	utils2 "github.com/Mir00r/user-service/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// AuthMiddleware validates the JWT and injects user claims into the Gin context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader(constants.Authorization)
		if authHeader == "" {
			utils2.GinErrorResponse(c, http.StatusUnauthorized, constants.ErrMissingAuthHeader)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, constants.Bearer)
		if tokenString == "" {
			utils2.GinErrorResponse(c, http.StatusUnauthorized, constants.ErrInvalidAuthHeader)
			c.Abort()
			return
		}

		// Verify the JWT and extract claims
		claims, err := utils2.VerifyJWT(tokenString)
		if err != nil {
			log.Printf("JWT verification error: %v", err)
			utils2.GinErrorResponse(c, http.StatusUnauthorized, constants.ErrInvalidToken)
			c.Abort()
			return
		}
		log.Printf("Extracted user ID: %s", claims.UserID)

		// Add claims to the Go context
		ctx := utils2.AddClaimsToContext(c.Request.Context(), claims)

		// Inject claims into the Gin context
		c.Set("userID", claims.UserID)
		c.Request = c.Request.WithContext(ctx) // Attach the modified context to the request

		log.Printf("Extracted user ID: %s", claims.UserID)
		c.Next()
	}
}

// BasicAuthMiddleware validates requests using Basic Auth or API keys
func BasicAuthMiddleware(c *gin.Context) {
	// Fetch expected credentials from environment variables
	expectedUsername := config.AppConfig.InternalSecurity.UserName
	expectedPassword := config.AppConfig.InternalSecurity.Password
	//expectedAPIKey := os.Getenv("API_KEY") // API key support

	// Check Basic Auth credentials
	username, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && username == expectedUsername && password == expectedPassword {
		c.Next()
		return
	}

	// Unauthorized if no valid credentials
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}
