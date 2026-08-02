package handler

import (
	"WhatShouldICook/internal/service"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	auth *service.AuthService
	user *service.UserService
}

func NewAuthHandler(auth *service.AuthService, user *service.UserService) *AuthHandler {
	return &AuthHandler{
		auth: auth,
		user: user,
	}
}

func (h *AuthHandler) Login(res http.ResponseWriter, req *http.Request) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.user.Login(req.Context(), request.Username, request.Password)
	if err != nil {
		http.Error(res, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := h.auth.GenerateToken(user)
	if err != nil {
		http.Error(res, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(map[string]string{"token": token})
}
