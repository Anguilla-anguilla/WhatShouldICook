package service

import "WhatShouldICook/internal/repository/postgres"

type CuisineService struct {
	repo *postgres.CuisineRepo
}