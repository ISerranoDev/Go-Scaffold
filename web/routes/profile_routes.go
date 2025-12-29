package routes

import (
	"github.com/go-chi/chi/v5"
	"go-scaffold.iserranodev.net/infrastructure/handler"
	"go-scaffold.iserranodev.net/infrastructure/middleware"
	"go-scaffold.iserranodev.net/internal/service"
	"os"
)

func RegisterProfileRoutes(r chi.Router, h *handler.ProfileHandler) {
	r.Route("/my-profile", func(r chi.Router) {
		jwtSvc := service.NewJWTService(os.Getenv("JWT_SECRET"))
		r.Use(middleware.JWTAuthMiddleware(jwtSvc))
		r.Get("/", h.Read)
		r.Put("/update", h.Update)
	})
}
