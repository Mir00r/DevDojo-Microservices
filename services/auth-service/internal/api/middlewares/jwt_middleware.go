package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"github.com/Mir00r/auth-service/internal/utils"
)

// AuthMiddleware validates the JWT and injects user claims into the Gin context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.GinErrorResponse(c, http.StatusUnauthorized, "Authorization header is missing")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			utils.GinErrorResponse(c, http.StatusUnauthorized, "Invalid Authorization header")
			c.Abort()
			return
		}

		// Verify the JWT and extract claims
		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			utils.GinErrorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		// Inject claims into the Gin context
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || !utils.ValidateBasicAuth(username, password) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
