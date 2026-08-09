package handler

import (
	"WhatShouldICook/internal/domain"
	"WhatShouldICook/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type RationHandler struct {
	service *service.RationService
}

func NewRationHandler(service *service.RationService) *RationHandler {
	return &RationHandler{service: service}
}

func (h *RationHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	idStr := req.PathValue("id")
	if idStr == "" {
		http.Error(res, "Missing id", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(res, "Invalid id", http.StatusBadRequest)
		return
	}

	userID, ok := req.Context().Value("userID").(int64)
	if !ok {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ration, err := h.service.GetByID(req.Context(), id, userID)
	if err != nil {
		if err == domain.ErrNotFound {
			http.Error(res, "Ration not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ration)
}
