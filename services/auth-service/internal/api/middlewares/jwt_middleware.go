package middlewares

import (
	"net/http"
	"strings"

	"github.com/Mir00r/auth-service/internal/utils"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}

		claims, err := utils.VerifyJWT(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Attach claims to the request context for further use
		ctx := utils.AddClaimsToContext(r.Context(), claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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
