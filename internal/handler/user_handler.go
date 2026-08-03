package handler

import (
	"WhatShouldICook/internal/domain"
	"WhatShouldICook/internal/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	id, ok := req.Context().Value("userID").(int64)
	if !ok {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.service.GetByID(req.Context(), id)
	if err != nil {
		if err == domain.ErrNotFound {
			http.Error(res, "User not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Internal error", http.StatusInternalServerError)
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(user)
}

func (h *UserHandler) UpdateProfile(res http.ResponseWriter, req http.Request) {
	id, ok := req.Context().Value("userID").(int64)
	if !ok {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
	}

	var request struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateProfile(req.Context(), id, request.UserName, request.Email); err != nil {
		switch err {
		case domain.ErrNotFound:
			http.Error(res, "User not found", http.StatusNotFound)
		case domain.ErrEmptyName:
			http.Error(res, "Name cannot be empty", http.StatusBadRequest)
		case domain.ErrInvalidEmail:
			http.Error(res, "Invalid email format", http.StatusBadRequest)
		default:
			http.Error(res, "Internal error", http.StatusInternalServerError)
		}
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(map[string]string{
		"message": "Profile updated",
	})
}

func (h *UserHandler) UpdatePassword(res http.ResponseWriter, req http.Request) {
	id, ok := req.Context().Value("userID").(int64)
	if !ok {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var request struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdatePassword(req.Context(), id, request.OldPassword, request.NewPassword); err != nil {
		switch err {
		case domain.ErrPasswordTooShort:
			http.Error(res, "Password should be 8 characters at least", http.StatusUnprocessableEntity)
		case domain.ErrInvalidPassword:
			http.Error(res, "Wrong password", http.StatusUnprocessableEntity)
		default:
			http.Error(res, "Internal error", http.StatusInternalServerError)
		}
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(map[string]string{
		"message": "Password updated",
	})
}

func (h *UserHandler) Delete(res http.ResponseWriter, req http.Request) {
	id, ok := req.Context().Value("userID").(int64)
	if !ok {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := h.service.Delete(req.Context(), id); err != nil {
		if err == domain.ErrNotFound {
			http.Error(res, "User not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusNoContent)
}
