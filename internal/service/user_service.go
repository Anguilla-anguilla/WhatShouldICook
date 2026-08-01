package service

import (
	"WhatShouldICook/internal/domain"
	"WhatShouldICook/internal/repository"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, username, email, password string) (*domain.User, error) {
	if len(password) < 8 {
		return nil, domain.ErrPasswordTooShort
	}

	_, err := s.repo.GetByUserName(ctx, username)
	if err != domain.ErrNotFound {
		return nil, domain.ErrAlreadyExists
	}
	if err != nil {
		return nil, err
	}

	_, err = s.repo.GetByEmail(ctx, email)
	if err != domain.ErrNotFound {
		return nil, domain.ErrAlreadyExists
	}
	if err != nil {
		return nil, err
	}

	hash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		UserName:     username,
		Email:        email,
		PasswordHash: hash,
	}
	err = user.Validate()
	if err != nil {
		return nil, err
	}

	if err = s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Вернуться сюда потом (наверное).
func (s *UserService) Login(ctx context.Context, username, password string) (*domain.User, error) {

	user, err := s.repo.GetByUserName(ctx, username)
	if err != nil {
		return nil, err
	}

	if err = comparePassword(password, user.PasswordHash); err != nil {
		return nil, domain.ErrInvalidPassword
	}

	return user, nil
}

func (s *UserService) GetByID(ctx context.Context, userID int64) (*domain.User, error) {
	return s.repo.GetByID(ctx, userID)
}

func (s *UserService) GetByUserName(ctx context.Context, userName string) (*domain.User, error) {
	return s.repo.GetByUserName(ctx, userName)
}

// Тут и в кухнях точно должна быть обертка?
func (s *UserService) Update(ctx context.Context, newUser *domain.User) error {
	if err := newUser.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ctx, newUser)
}

func (s *UserService) UpdatePassword(ctx context.Context, user *domain.User, oldPassword, newPassword string) error {
	if err := comparePassword(oldPassword, user.PasswordHash); err != nil {
		return domain.ErrInvalidPassword
	}

	if len(newPassword) < 8 {
		return domain.ErrPasswordTooShort
	}

	newHash, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	newPassUser := &domain.User{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: newHash,
	}

	if err = s.Update(ctx, newPassUser); err != nil {
		return err
	}
	return nil
}

func (s *UserService) Delete(ctx context.Context, user *domain.User) error {
	return s.repo.Delete(ctx, user.ID)
}

func hashPassword(password string) (hash string, err error) {
	crypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(crypt), nil
}

func comparePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
