package routes

import (
	"api_gateway/internal/config"
	"api_gateway/internal/transport/http/middleware"
	"github.com/go-chi/chi/v5"
)

func Init(conf *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.ApiMiddleware())

	router.Route("/api/users", func(r chi.Router) {
		UserRoutes(r, &conf.AuthService)
	})
	return router
}
