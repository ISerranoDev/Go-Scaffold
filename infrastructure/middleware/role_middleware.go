package middleware

import (
	"go-react-backend.iserranodev.net/internal/helper"
	"net/http"
	"strings"
)

func RoleAuthorizationMiddleware(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			roles, ok := r.Context().Value("roles").([]string)
			if !ok || len(roles) == 0 {
				helper.RespondJSON(w, http.StatusForbidden, false, nil, []string{"Acceso denegado: no autenticado o sin roles"})
				return
			}

			for _, role := range roles {
				if strings.EqualFold(role, requiredRole) {
					next.ServeHTTP(w, r)
					return
				}
			}

			helper.RespondJSON(w, http.StatusForbidden, false, nil, []string{"Acceso denegado: rol insuficiente"})
		})
	}
}
