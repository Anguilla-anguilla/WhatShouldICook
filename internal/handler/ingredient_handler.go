package handler

import (
	"WhatShouldICook/internal/domain"
	"WhatShouldICook/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type IngredientHandler struct {
	service *service.IngredientService
}

func NewIngredientHandler(service *service.IngredientService) *IngredientHandler {
	return &IngredientHandler{service: service}
}

func (h *IngredientHandler) Create(res http.ResponseWriter, req *http.Request) {
	var request struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	ingredient, err := h.service.Create(req.Context(), request.Name)
	if err != nil {
		if err == domain.ErrEmptyName {
			http.Error(res, "Name cannot be empty", http.StatusBadRequest)
			return
		}
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(ingredient)
}

func (h *IngredientHandler) List(res http.ResponseWriter, req *http.Request) {
	ingredients, err := h.service.List(req.Context())
	if err != nil {
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ingredients)
}

func (h *IngredientHandler) GetByID(res http.ResponseWriter, req *http.Request) {
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

	ingredient, err := h.service.GetByID(req.Context(), id)
	if err != nil {
		switch err {
		case domain.ErrNotFound:
			http.Error(res, "Ingredient not found", http.StatusNotFound)
		case domain.ErrEmptyName:
			http.Error(res, "Name cannot be empty", http.StatusBadRequest)
		default:
			http.Error(res, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ingredient)
}

// func (hh IngredientHandler) GetByName() {}
func (h *IngredientHandler) Update(res http.ResponseWriter, req *http.Request) {
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

	var request struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(req.Context(), id, request.Name); err != nil {
		switch err {
		case domain.ErrNotFound:
			http.Error(res, "Ingredient not found", http.StatusNotFound)
		case domain.ErrEmptyName:
			http.Error(res, "Name cannot be empty", http.StatusBadRequest)
		default:
			http.Error(res, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(map[string]string{"message": "Updated"})
}

func (h *IngredientHandler) Delete(res http.ResponseWriter, req *http.Request) {
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

	if err := h.service.Delete(req.Context(), id); err != nil {
		if err == domain.ErrNotFound {
			http.Error(res, "Ingredient not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusNoContent)
}
