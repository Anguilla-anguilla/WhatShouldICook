package service

import (
	"WhatShouldICook/internal/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	secret string
	ttl    time.Duration
	user   *UserService
}

func NewAuthService(secret string, ttl int, userService *UserService) *AuthService {
	return &AuthService{
		secret: secret,
		ttl:    time.Duration(ttl) * time.Hour,
		user:   userService,
	}
}

func (a *AuthService) GenerateToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.UserName,
		"exp":      time.Now().Add(a.ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(a.secret))
}

func (a *AuthService) ValidateToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(a.secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0, errors.New("invalid user_id in token")
		}
		return int64(userIDFloat), nil
	}
	return 0, errors.New("invalid token")
}
