package service

import "WhatShouldICook/internal/domain"

type GenerateMenuRequest struct {
	UserID        int64
	CuisineID     int64
	CategoryCount map[int64]int
	Duration      int64
}

type MenuResponse struct {
	Ration       *domain.Ration
	Recipes      []*domain.Recipe
	ShoppingList []ShoppingItem
}

type ShoppingItem struct {
	IngredientName string `json:"ingredient_name"`
	TotalQuantity  int64  `json:"total_quantity"`
}
