package routes

import (
	"github.com/go-chi/chi/v5"
	"go-react-backend.iserranodev.net/infrastructure/handler"
)

func RegisterAuthRoutes(r chi.Router, h *handler.AuthHandler) {
	r.Post("/login", h.Login)
}
