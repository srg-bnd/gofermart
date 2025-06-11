package services

import (
	"context"
	"errors"
	"ya41-56/internal/gophermart/models"
	"ya41-56/internal/gophermart/repositories"
	"ya41-56/internal/shared/bcryptutil"
)

type AuthService struct {
	Users repositories.Repository[models.User]
}

var ErrInvalidCreds = errors.New("invalid credentials")

func NewAuthService(users repositories.Repository[models.User]) *AuthService {
	return &AuthService{Users: users}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (access string, refresh string, err error) {
	user, err := s.Users.FindByField(ctx, "email", email)
	if err != nil {
		return "", "", ErrInvalidCreds
	}

	if err := bcryptutil.CheckHash(password, user.PasswordHash); err != nil {
		return "", "", ErrInvalidCreds
	}

	return "", "", nil // TODO: тут твоя реализация
}

func (s *AuthService) Register(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := s.Users.FindByField(ctx, "email", user.Login)
	if err == nil {
		return user, errors.New("user already exists")
	}

	hashed, err := bcryptutil.Hash(user.Password)
	if err != nil {
		return user, err
	}

	user.PasswordHash = hashed

	if err := s.Users.Create(ctx, user); err != nil {
		return user, err
	}
	return user, nil
}

func (s *AuthService) ParseAndValidate(_ string) (models.User, error) {
	return models.User{}, nil // TODO: тут твоя реализация
}
