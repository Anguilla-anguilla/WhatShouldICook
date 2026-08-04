package handler

import (
	"WhatShouldICook/internal/domain"
	"WhatShouldICook/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) List(res http.ResponseWriter, req *http.Request) {
	categories, err := h.service.List(req.Context())
	if err != nil {
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(categories)
}

func (h *CategoryHandler) GetByID(res http.ResponseWriter, req *http.Request) {
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

	category, err := h.service.GetByID(req.Context(), id)
	if err != nil {
		if err == domain.ErrNotFound {
			http.Error(res, "Category not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Internal Error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(category)
}

func (h *CategoryHandler) GetByName(res http.ResponseWriter, req *http.Request) {
	name := req.PathValue("name")
	if name == "" {
		http.Error(res, "Missing name", http.StatusBadRequest)
		return
	}

	category, err := h.service.GetByName(req.Context(), name)
	if err != nil {
		if err == domain.ErrNotFound {
			http.Error(res, "Category not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(category)
}
