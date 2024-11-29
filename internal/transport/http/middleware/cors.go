package middleware

import (
	"github.com/go-chi/cors"
	"net/http"
)

func CorsMiddleware() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	})
}
