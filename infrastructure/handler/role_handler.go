package handler

import (
	"go-react-backend.iserranodev.net/internal/dto"
	"go-react-backend.iserranodev.net/internal/helper"
	"go-react-backend.iserranodev.net/internal/usecase"
	"net/http"
)

type RoleHandler struct {
	rc *usecase.RoleUseCase
}

func NewRoleHandler(rc *usecase.RoleUseCase) *RoleHandler {
	return &RoleHandler{rc}
}

func (h *RoleHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	roles, err := h.rc.GetAll()
	if err != nil {
		helper.RespondJSON(w, http.StatusInternalServerError, false, nil, []string{err.Error()})
		return
	}

	var resp []dto.RoleResponse
	for _, r := range roles {
		resp = append(resp, dto.FromDomainRole(r))
	}

	helper.RespondJSON(w, http.StatusOK, true, resp, nil)
}
