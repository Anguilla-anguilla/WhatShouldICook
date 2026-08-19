package handler

import (
	"WhatShouldICook/internal/domain"
	"WhatShouldICook/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type RecipeHandler struct {
	service *service.RecipeService
}

func NewRecipeHandler(service *service.RecipeService) *RecipeHandler {
	return &RecipeHandler{service: service}
}

// ДОБАВИТЬ БОЛЬШЕ ОБРАБОТОК ОШИБОК
// заменить на errors.Is()
func (h *RecipeHandler) Create(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value("userID").(int64)
	if !ok {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var request service.CreateRecipeRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	recipe, err := h.service.Create(req.Context(), request, userID)
	if err != nil {
		switch err {
		case domain.ErrEmptyName:
			http.Error(res, "Name is required", http.StatusBadRequest)
		default:
			http.Error(res, "Internal error", http.StatusInternalServerError)
		}
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(recipe)
}

func (h *RecipeHandler) List(res http.ResponseWriter, req *http.Request) {

	filters := service.RecipeFilters{}

	userID, ok := req.Context().Value("userID").(int64)
	if !ok {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}
	filters.UserID = userID

	if categoryID := req.URL.Query().Get("category_id"); categoryID != "" {
		id, err := strconv.ParseInt(categoryID, 10, 64)
		if err != nil {
			http.Error(res, "Wrong params", http.StatusBadRequest)
			return
		}
		filters.CategoryID = &id
	}

	if cuisineID := req.URL.Query().Get("cuisine_id"); cuisineID != "" {
		id, err := strconv.ParseInt(cuisineID, 10, 64)
		if err != nil {
			http.Error(res, "Wrong params", http.StatusBadRequest)
			return
		}
		filters.CuisineID = &id
	}

	if favorite := req.URL.Query().Get("favorite"); favorite != "" {
		fav, err := strconv.ParseBool(favorite)
		if err != nil {
			http.Error(res, "Wrong params", http.StatusBadRequest)
			return
		}
		filters.Favorite = &fav
	}

	if public := req.URL.Query().Get("public"); public != "" {
		pub, err := strconv.ParseBool(public)
		if err != nil {
			http.Error(res, "Wrong params", http.StatusBadRequest)
			return
		}
		filters.Public = &pub
	}

	recipes, err := h.service.List(req.Context(), filters)
	if err != nil {
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(recipes)
}

func (h *RecipeHandler) GetByID(res http.ResponseWriter, req *http.Request) {
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

	recipe, err := h.service.GetByID(req.Context(), id, userID)
	if err != nil {
		if err == domain.ErrNotFound {
			http.Error(res, "Recipe not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(recipe)
}

func (h *RecipeHandler) Update(res http.ResponseWriter, req *http.Request) {
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

	var request service.CreateRecipeRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(req.Context(), request, id, userID); err != nil {
		switch err {
		case domain.ErrNotFound:
			http.Error(res, "Recipe not found", http.StatusNotFound)
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

func (h *RecipeHandler) Delete(res http.ResponseWriter, req *http.Request) {
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

	if err := h.service.Delete(req.Context(), id, userID); err != nil {
		if err == domain.ErrNotFound {
			http.Error(res, "Recipe not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Internal error", http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusNoContent)
}

func (h *RecipeHandler) Copy(res http.ResponseWriter, req *http.Request) {
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

	var request struct {
		CuisineID int64 `json:"cuisine_id"`
		OwnerID   int64 `json:"owner_id"`
	}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	recipe, err := h.service.Copy(req.Context(), id, userID, request.OwnerID, request.CuisineID)
	if err != nil {
		switch err {
		case domain.ErrPermissionDenied:
			http.Error(res, "Recipe is not public", http.StatusForbidden)
		case domain.ErrNotFound:
			http.Error(res, "Recipe not found", http.StatusNotFound)
		default:
			http.Error(res, "Internal error", http.StatusInternalServerError)
		}
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(recipe)
}
