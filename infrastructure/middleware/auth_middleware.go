package middleware

import (
	"context"
	"go-scaffold.iserranodev.net/internal/helper"
	"go-scaffold.iserranodev.net/internal/service"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(jwtSvc *service.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
				helper.RespondJSON(w, http.StatusForbidden, false, nil, []string{"No se ha proporcionado la cabecera: Authorization"})
				return
			}

			token := strings.TrimPrefix(auth, "Bearer ")
			claims, err := jwtSvc.ValidateToken(token)
			if err != nil {
				helper.RespondJSON(w, http.StatusForbidden, false, nil, []string{"Token Inválido"})
				return
			}

			ctx := context.WithValue(r.Context(), "roles", claims.Roles)
			ctx = context.WithValue(ctx, "user_id", claims.UserId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
