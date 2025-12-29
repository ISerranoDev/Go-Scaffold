package handler

import (
	"encoding/json"
	"go-scaffold.iserranodev.net/internal/dto"
	"go-scaffold.iserranodev.net/internal/helper"
	"go-scaffold.iserranodev.net/internal/service"
	"go-scaffold.iserranodev.net/internal/usecase"
	"net/http"
)

type AuthHandler struct {
	uc  *usecase.AuthUseCase
	jwt *service.JWTService
}

func NewAuthHandler(uc *usecase.AuthUseCase, jwt *service.JWTService) *AuthHandler {
	return &AuthHandler{uc, jwt}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.RespondJSON(w, http.StatusBadRequest, false, nil, []string{"Cuerpo del mensaje no válido"})
		return
	}

	user, err := h.uc.Authenticate(req.Email, req.Password)
	if err != nil {
		helper.RespondJSON(w, http.StatusUnauthorized, false, nil, []string{"Credenciales no válidas"})
		return
	}

	token, err := h.jwt.GenerateToken(user)
	if err != nil {
		helper.RespondJSON(w, http.StatusInternalServerError, false, nil, []string{"Error generating token"})
		return
	}

	userResponse := dto.FromDomainUser(user)

	helper.RespondJSON(w, http.StatusOK, true, map[string]any{"token": token, "user": userResponse}, nil)
}
