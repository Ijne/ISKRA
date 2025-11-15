package middleware

import (
	"context"
	"log"
	"net/http"
	"strconv"
)

func UserMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
			if err != nil {
				log.Println(err)
			}

			ctx := context.WithValue(r.Context(), UserIDKey, id)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
