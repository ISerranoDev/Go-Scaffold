package routes

import (
	"github.com/go-chi/chi/v5"
	"go-react-backend.iserranodev.net/infrastructure/handler"
	"go-react-backend.iserranodev.net/infrastructure/middleware"
	"go-react-backend.iserranodev.net/internal/service"
	"os"
)

func RegisterUserRoutes(r chi.Router, h *handler.UserHandler) {
	r.Route("/users", func(r chi.Router) {
		jwtSvc := service.NewJWTService(os.Getenv("JWT_SECRET"))
		r.Use(middleware.JWTAuthMiddleware(jwtSvc))
		r.Use(middleware.RoleAuthorizationMiddleware("admin"))
		r.Get("/", h.List)
		r.Post("/create", h.Create)
		r.Get("/read", h.Read)
		r.Put("/update", h.Update)
		r.Delete("/delete", h.Delete)
	})
}
