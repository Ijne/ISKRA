package middleware

import (
	"context"
	"iskra/miniapp/internal/tools/auth"
	"iskra/shared/config"
	"net/http"
	"strings"
)

const (
	UserIDKey = "userID"
)

func JWTAuthMiddleware(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			// формат заголовка
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]

			// парсим и проверяем токен
			userID, err := auth.GetDataFromJWTToken(tokenString, cfg)
			switch err {
			case auth.ErrInvalidToken:
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			case auth.ErrTokenExpired:
				http.Error(w, "Token expired", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// заглушка
// func GetUserIDFromContext(ctx context.Context) (int64, bool) {
// 	return 1, true
// }

// реальный код
func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey).(int64)
	return userID, ok
}
