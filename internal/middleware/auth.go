package middleware

import (
	"WhatShouldICook/internal/service"
	"context"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	auth *service.AuthService
}

func NewAuthMiddleware(auth *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{auth: auth}
}

func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(res, "Missing auth header", http.StatusUnauthorized)
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(res, "Invalid header format", http.StatusUnauthorized)
			return
		}
		token := parts[1]
		userID, err := m.auth.ValidateToken(token)
		if err != nil {
			http.Error(res, "Invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(req.Context(), "userID", userID)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
