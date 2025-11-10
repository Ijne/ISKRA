package middleware

import (
	"net/http"
)

func CorsMiddleware(allowOrigin string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			next.ServeHTTP(w, r)
		})
	}
}
