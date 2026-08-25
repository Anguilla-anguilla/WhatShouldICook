package service

import (
	"WhatShouldICook/internal/domain"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, username, email, password string) (*domain.User, error) {
	if len(password) < 8 {
		return nil, domain.ErrPasswordTooShort
	}

	_, err := s.repo.GetByUserName(ctx, username)
	if err != nil && err != domain.ErrNotFound {
		return nil, err
	}
	if err == nil {
		return nil, domain.ErrAlreadyExists
	}

	_, err = s.repo.GetByEmail(ctx, email)
	if err != nil && err != domain.ErrNotFound {
		return nil, err
	}
	if err == nil {
		return nil, domain.ErrAlreadyExists
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

func (s *UserService) GetByID(ctx context.Context, userID int64) (*UserResponse, error) {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	responce := UserResponse{ID: user.ID, Name: user.UserName, Email: user.Email}
	return &responce, nil
}

func (s *UserService) GetByUserName(ctx context.Context, userName string) (*UserResponse, error) {
	user, err := s.repo.GetByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}
	responce := UserResponse{ID: user.ID, Name: user.UserName, Email: user.Email}
	return &responce, nil
}

// не уверена, что оно будет работать. Мэй би, сделать, как в пассворде? хотя хз.
func (s *UserService) UpdateProfile(ctx context.Context, id int64, username, email string) error {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	user.UserName = username
	user.Email = email

	if err := user.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ctx, user)
}

func (s *UserService) UpdatePassword(ctx context.Context, id int64, oldPassword, newPassword string) error {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

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

	user.PasswordHash = newHash

	if err = s.repo.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
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
