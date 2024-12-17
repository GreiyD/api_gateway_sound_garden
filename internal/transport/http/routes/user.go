package routes

import (
	"api_gateway/internal/config"
	"api_gateway/internal/handlers/auth"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func UserRoutes(router chi.Router, conf *config.AuthService) {
	router.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		auth.RegisterUser(w, r, conf)
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.LoginUser(w, r, conf)
	})

	router.Post("/logout", func(w http.ResponseWriter, r *http.Request) {
		auth.LogoutUser(w, r, conf)
	})
}
