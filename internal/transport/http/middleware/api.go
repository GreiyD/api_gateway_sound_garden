package middleware

import (
	"net/http"
	"strings"
)

func ApiMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Header.Get("Content-Type") != "application/json" {
				http.Error(w, "Invalid data format. Expected application/json", http.StatusBadRequest)
				return
			}

			if !strings.HasPrefix(r.URL.Path, "/api/") {
				http.Error(w, "Invalid path, must start with /api", http.StatusBadRequest)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
