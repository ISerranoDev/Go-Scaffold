package routes

import (
	"github.com/go-chi/chi/v5"
	"go-scaffold.iserranodev.net/infrastructure/handler"
	"go-scaffold.iserranodev.net/infrastructure/middleware"
	"go-scaffold.iserranodev.net/internal/service"
	"os"
)

func RegisterRoleRoutes(r chi.Router, h *handler.RoleHandler) {
	r.Route("/roles", func(r chi.Router) {
		jwtSvc := service.NewJWTService(os.Getenv("JWT_SECRET"))
		r.Use(middleware.JWTAuthMiddleware(jwtSvc))
		r.Use(middleware.RoleAuthorizationMiddleware("admin"))
		r.Get("/", h.GetAll)
	})
}
