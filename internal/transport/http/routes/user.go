package routes

import (
	"api_gateway/internal/handlers/auth"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(router chi.Router) {
	router.Get("/", auth.GetUsers)
	router.Post("/register", auth.RegisterUser)
}
