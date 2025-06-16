package services

import (
	"context"
	"ya41-56/internal/gophermart/customerror"
	"ya41-56/internal/gophermart/models"
	"ya41-56/internal/shared/bcryptutil"
	"ya41-56/internal/shared/repositories"
)

type AuthService struct {
	Users        repositories.Repository[models.User]
	TokenService *TokenService
}

func NewAuthService(users repositories.Repository[models.User], tokenService *TokenService) *AuthService {
	return &AuthService{
		Users:        users,
		TokenService: tokenService,
	}
}

func (s *AuthService) Login(ctx context.Context, login, password string) (string, error) {
	user, err := s.Users.FindByField(ctx, "login", login)
	if err != nil {
		return "", customerror.ErrInvalidCreds
	}

	if err = bcryptutil.CheckHash(password, user.PasswordHash); err != nil {
		return "", customerror.ErrInvalidCreds
	}

	if user.Status == models.UserStatusDisabled {
		return "", customerror.ErrInvalidCreds
	}

	return s.TokenService.BuildJWTString(user.ID)
}

func (s *AuthService) ParseAndValidate(_ context.Context, tokenString string) (string, error) {
	claims := Claims{}
	token, err := s.TokenService.ParseToken(&claims, tokenString)
	if err != nil || !token.Valid {
		return "", customerror.ErrJWTToken
	}

	return claims.UserID, nil
}

func (s *AuthService) Register(ctx context.Context, user *models.User) (string, error) {
	_, err := s.Users.FindByField(ctx, "login", user.Login)
	if err == nil {
		return "", customerror.ErrUserExists
	}

	hashed, err := bcryptutil.Hash(user.Password)
	if err != nil {
		return "", err
	}

	user.PasswordHash = hashed
	user.Status = models.UserStatusActive

	if err := s.Users.Create(ctx, user); err != nil {
		return "", err
	}

	return s.TokenService.BuildJWTString(user.ID)
}
