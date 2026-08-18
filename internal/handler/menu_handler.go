package handler

import (
	"WhatShouldICook/internal/service"
	"encoding/json"
	"net/http"
)

type MenuHandler struct {
	service *service.MenuService
}

func NewMenuHandler(service *service.MenuService) *MenuHandler {
	return &MenuHandler{service: service}
}

func (h *MenuHandler) Generate(res http.ResponseWriter, req http.Request) {
	userID, ok := req.Context().Value("userID").(int64)
	if !ok {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
	}

	var request struct {
		CuisineID     int64         `json:"cuisine_id"`
		CategoryCount map[int64]int `json:"category_count"`
		Duration      int64         `json:"duration"`
	}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}
	serviceReq := service.GenerateMenuRequest{
		UserID:        userID,
		CuisineID:     request.CuisineID,
		CategoryCount: request.CategoryCount,
		Duration:      request.Duration,
	}

	menu, err := h.service.GenerateMenu(req.Context(), serviceReq)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(menu)
}
