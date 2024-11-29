package routes

import (
	"api_gateway/internal/transport/http/middleware"
	"github.com/go-chi/chi/v5"
)

func Init() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.ApiMiddleware())

	router.Route("/api/users", UserRoutes)

	return router
}
