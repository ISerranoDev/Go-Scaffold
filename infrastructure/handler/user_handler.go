package handler

import (
	"encoding/json"
	"go-react-backend.iserranodev.net/internal/dto"
	"go-react-backend.iserranodev.net/internal/helper"
	"go-react-backend.iserranodev.net/internal/usecase"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	uc *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	email := query.Get("email")
	rolesStr := query.Get("roles") // ej: "1,2,3"
	pageStr := query.Get("page")
	pageSizeStr := query.Get("pageSize")
	sortStr := query.Get("sort") // ej: "username:ASC,email:DESC"

	var roles []int
	if rolesStr != "" {
		for _, rp := range strings.Split(rolesStr, ",") {
			if id, err := strconv.Atoi(rp); err == nil {
				roles = append(roles, id)
			}
		}
	}

	page := 1
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	pageSize := 20
	if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
		pageSize = ps
	}

	sortMap := make(map[string]string)
	if sortStr != "" {
		for _, s := range strings.Split(sortStr, ",") {
			parts := strings.Split(s, ":")
			if len(parts) == 2 {
				field := parts[0]
				order := strings.ToUpper(parts[1])
				if order != "ASC" && order != "DESC" {
					order = "ASC"
				}
				sortMap[field] = order
			}
		}
	}

	filters := dto.UserFilter{
		Filters: dto.FilterFields{
			Email: email,
			Roles: roles,
		},
		Page:     page,
		PageSize: pageSize,
		Sort:     sortMap,
	}

	users, counter, err := h.uc.List(filters)
	if err != nil {
		helper.RespondJSON(w, http.StatusInternalServerError, false, nil, []string{err.Error()})
		return
	}

	var userDtos []dto.UserResponse
	for _, u := range users {
		userDtos = append(userDtos, dto.FromDomainUser(u))
	}

	// Estructura de respuesta
	resp := map[string]interface{}{
		"items": userDtos,
		"total": counter,
	}

	helper.RespondJSON(w, http.StatusOK, true, resp, nil)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.RespondJSON(w, http.StatusBadRequest, false, nil, []string{"Datos inválidos"})
		return
	}

	if input.Password != input.RePassword {
		helper.RespondJSON(w, http.StatusBadRequest, false, nil, []string{"Las contraseñas no coinciden"})
		return
	}

	user, err := h.uc.Create(&input)
	if err != nil {
		helper.RespondJSON(w, http.StatusInternalServerError, false, nil, []string{err.Error()})
		return
	}

	helper.RespondJSON(w, http.StatusCreated, true, dto.FromDomainUser(user), nil)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.RespondJSON(w, http.StatusBadRequest, false, nil, []string{"Datos inválidos"})
		return
	}

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

func (h *UserHandler) Read(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		helper.RespondJSON(w, http.StatusBadRequest, false, nil, []string{"ID del usuario requerido"})
		return
	}

	user, err := h.uc.GetByID(id)
	if err != nil {
		helper.RespondJSON(w, http.StatusNotFound, false, nil, []string{err.Error()})
		return
	}

	helper.RespondJSON(w, http.StatusOK, true, dto.FromDomainUser(user), nil)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		helper.RespondJSON(w, http.StatusBadRequest, false, nil, []string{"ID del usuario requerido"})
		return
	}

	if err := h.uc.Delete(id); err != nil {
		helper.RespondJSON(w, http.StatusInternalServerError, false, nil, []string{err.Error()})
		return
	}

	helper.RespondJSON(w, http.StatusOK, true, "Usuario eliminado correctamente", nil)
}
