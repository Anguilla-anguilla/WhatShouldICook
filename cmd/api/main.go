package main

import (
	"WhatShouldICook/internal/config"
	"WhatShouldICook/internal/handler"
	midd "WhatShouldICook/internal/middleware"
	"WhatShouldICook/internal/repository/postgres"
	"WhatShouldICook/internal/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.Database.URL)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping database %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // при панике возвращает 500 не роняя сервер
	r.Use(middleware.Timeout(60 * time.Second))

	userRepo := postgres.NewUserRepo(pool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	authService := service.NewAuthService(cfg.JWT.Secret, cfg.JWT.TTL)
	authMiddleware := midd.NewAuthMiddleware(authService)
	authHandler := handler.NewAuthHandler(authService, userService)
	r.Route("/api/v1/users", func(r chi.Router) {
		r.Use(authMiddleware.RequireAuth)
		r.Get("/profile", userHandler.GetByID)
		r.Put("/profile", userHandler.UpdateProfile)
		r.Put("/password", userHandler.UpdatePassword)
		r.Delete("/profile", userHandler.Delete)
	})
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
	})

	categoryRepo := postgres.NewCategoryRepo(pool)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	r.Route("/api/v1/categories", func(r chi.Router) {
		r.Use(authMiddleware.RequireAuth)
		r.Get("/", categoryHandler.List)
		r.Get("/{id}", categoryHandler.GetByID)
	})

	cuisineRepo := postgres.NewCuisineRepo(pool)
	cuisineService := service.NewCuisineService(cuisineRepo)
	cuisineHandler := handler.NewCuisineHandler(cuisineService)
	r.Route("/api/v1/cuisines", func(r chi.Router) {
		r.Use(authMiddleware.RequireAuth)
		r.Get("/", cuisineHandler.List)
		r.Post("/", cuisineHandler.Create)
		r.Get("/{id}", cuisineHandler.GetByID)
		r.Put("/{id}", cuisineHandler.Update)
		r.Delete("/{id}", cuisineHandler.Delete)
	})

	ingredientRepo := postgres.NewIngredientRepo(pool)
	ingredientService := service.NewIngredientService(ingredientRepo)
	ingredientHandler := handler.NewIngredientHandler(ingredientService)
	r.Route("/api/v1/ingredients", func(r chi.Router) {
		r.Use(authMiddleware.RequireAuth)
		r.Get("/", ingredientHandler.List)
		r.Post("/", ingredientHandler.Create)
		r.Get("/{id}", ingredientHandler.GetByID)
		r.Put("/{id}", ingredientHandler.Update)
		r.Delete("/{id}", ingredientHandler.Delete)
	})

	recipeIngredientRepo := postgres.NewRecipeIngredientRepo(pool)

	recipeRepo := postgres.NewRecipeRepo(pool)
	recipeService := service.NewRecipeService(
		recipeRepo,
		recipeIngredientRepo,
		ingredientService,
		cuisineService,
		categoryService,
	)
	recipeHandler := handler.NewRecipeHandler(recipeService)
	r.Route("/api/v1/recipes", func(r chi.Router) {
		r.Use(authMiddleware.RequireAuth)
		r.Get("/", recipeHandler.List)
		r.Post("/", recipeHandler.Create)
		r.Get("/{id}", recipeHandler.GetByID)
		r.Put("/{id}", recipeHandler.Update)
		r.Delete("/{id}", recipeHandler.Delete)
		r.Post("/{id}/copy", recipeHandler.Copy)
	})

	shoppingListRepo := postgres.NewShoppingListRepo(pool)
	shoppingListService := service.NewShoppingListService(shoppingListRepo)
	shoppingListHandler := handler.NewShoppingListHandler(shoppingListService)
	r.Route("/api/v1/shopping-list", func(r chi.Router) {
		r.Use(authMiddleware.RequireAuth)
		r.Get("/{id}", shoppingListHandler.GetByID)
	})

	rationRecipeRepo := postgres.NewRationRecipeRepo(pool)
	rationRepo := postgres.NewRationRepo(pool)
	rationService := service.NewRationService(rationRepo, rationRecipeRepo, recipeService)
	rationHandler := handler.NewRationHandler(rationService)
	r.Route("/api/v1/ration", func(r chi.Router) {
		r.Use(authMiddleware.RequireAuth)
		r.Get("/{id}", rationHandler.GetByID)
	})

	shoppingListRecipeRepo := postgres.NewShoppingListRecipeRepo(pool)
	menuService := service.NewMenuService(
		recipeRepo,
		recipeIngredientRepo,
		rationRepo,
		rationRecipeRepo,
		shoppingListRepo,
		shoppingListRecipeRepo,
		ingredientRepo,
	)
	menuHandler := handler.NewMenuHandler(menuService)
	r.Route("/api/v1/menu", func(r chi.Router) {
		r.Use(authMiddleware.RequireAuth)
		r.Post("/generate", menuHandler.Generate)
	})

	srv := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Starting server on addr: %v", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server")
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server shutdown failed: %v.\nIt will work forever(", err)
	}

	log.Println("Server stopped")

}
