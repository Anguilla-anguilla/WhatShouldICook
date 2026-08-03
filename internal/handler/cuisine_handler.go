package handler

import (
	"WhatShouldICook/internal/domain"
	"WhatShouldICook/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type CuisineHandler struct {
	service *service.CuisineService
}

func NewCuisineHandler(service *service.CuisineService) *CuisineHandler {
	return &CuisineHandler{service: service}
}

func (c *CuisineHandler) Create(res http.ResponseWriter, req *http.Request) {
	var request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	userID, ok := req.Context().Value("userID").(int64)
	if !ok {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	cuisine, err := c.service.Create(req.Context(), request.Name, request.Description, userID)
	if err != nil {
		switch err {
		case domain.ErrEmptyName:
			http.Error(res, "Name cannot be empty", http.StatusBadRequest)
		case domain.ErrAlreadyExists:
			http.Error(res, "Cuisine already exists", http.StatusConflict)
		default:
			http.Error(res, "Internal error", http.StatusInternalServerError)
		}
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(cuisine)
}

// мб тут тоже гет бай нэйм нужна? надо подумотьб..
func (c *CuisineHandler) GetByID(res http.ResponseWriter, req *http.Request) {
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

	cuisine, err := c.service.GetByID(req.Context(), id, userID)
	if err != nil {
		if err == domain.ErrNotFound {
			http.Error(res, "Cuisine not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(cuisine)
}

func (c *CuisineHandler) List(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value("userID").(int64)
	if !ok {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	cuisines, err := c.service.List(req.Context(), userID)
	if err != nil {
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(cuisines)
}

func (c *CuisineHandler) Update(res http.ResponseWriter, req *http.Request) {
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
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	userID, ok := req.Context().Value("userID").(int64)
	if !ok {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	cuisine := &domain.Cuisine{
		ID:          id,
		Name:        request.Name,
		Description: request.Description,
	}
	// Повторяющееся можно вынести в утилиты или отдельные функции
	if err := c.service.Update(req.Context(), cuisine, userID); err != nil {
		switch err {
		case domain.ErrNotFound:
			http.Error(res, "Cuisine not found", http.StatusNotFound)
		case domain.ErrEmptyName:
			http.Error(res, "Name cannot be empty", http.StatusBadRequest)
		default:
			http.Error(res, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(cuisine)
}

func (c *CuisineHandler) Delete(res http.ResponseWriter, req *http.Request) {
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

	if err := c.service.Delete(req.Context(), id, userID); err != nil {
		if err == domain.ErrNotFound {
			http.Error(res, "Cuisine not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusNoContent)
}
