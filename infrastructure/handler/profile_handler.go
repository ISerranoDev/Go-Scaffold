package handler

import (
	"encoding/json"
	"go-react-backend.iserranodev.net/internal/dto"
	"go-react-backend.iserranodev.net/internal/helper"
	"go-react-backend.iserranodev.net/internal/usecase"
	"net/http"
)

type ProfileHandler struct {
	uc *usecase.ProfileUseCase
}

func NewProfileHandler(uc *usecase.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{uc}
}

func (h *ProfileHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.RespondJSON(w, http.StatusBadRequest, false, nil, []string{"Datos inválidos"})
		return
	}
	input.ID = r.Context().Value("user_id").(string)

	if (input.Password != "" || input.RePassword != "") && input.Password != input.RePassword {
		helper.RespondJSON(w, http.StatusBadRequest, false, nil, []string{"Las contraseñas no coinciden"})
		return
	}

	user, err := h.uc.Update(&input)
	if err != nil {
		helper.RespondJSON(w, http.StatusInternalServerError, false, nil, []string{err.Error()})
		return
	}

	helper.RespondJSON(w, http.StatusOK, true, dto.FromDomainUser(user), nil)
}

func (h *ProfileHandler) Read(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user_id").(string)
	if id == "" {
		helper.RespondJSON(w, http.StatusBadRequest, false, nil, []string{"ID de usuario inválido"})
		return
	}

	user, err := h.uc.GetByID(id)
	if err != nil {
		helper.RespondJSON(w, http.StatusNotFound, false, nil, []string{err.Error()})
		return
	}

	helper.RespondJSON(w, http.StatusOK, true, dto.FromDomainUser(user), nil)
}
