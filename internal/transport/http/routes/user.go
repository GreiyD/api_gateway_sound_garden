package routes

import (
	"api_gateway/internal/config"
	"api_gateway/internal/handlers/auth"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func UserRoutes(router chi.Router, conf *config.AuthService) {
	router.Get("/", auth.GetUsers)

	router.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		auth.RegisterUser(w, r, conf)
	})
}
