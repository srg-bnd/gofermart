package services

import (
	"context"
	"errors"
	"fmt"
	"time"
	"ya41-56/internal/gophermart/models"
	"ya41-56/internal/shared/bcryptutil"
	"ya41-56/internal/shared/repositories"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	Users     repositories.Repository[models.User]
	secretKey string
}

type Claims struct {
	jwt.RegisteredClaims
	UserLogin string `json:"login"`
}

const DEFAULT_TOKEN_EXP = time.Hour * 1

var ErrInvalidCreds = errors.New("invalid credentials")
var ErrUserExists = errors.New("user already exists")
var ErrJWTToken = errors.New("invalid JWT token")

func NewAuthService(users repositories.Repository[models.User], secretKey string) *AuthService {
	return &AuthService{
		Users:     users,
		secretKey: secretKey,
	}
}

func (s *AuthService) BuildJWTString(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(DEFAULT_TOKEN_EXP)),
		},
		UserLogin: login,
	})

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", ErrJWTToken
	}

	return tokenString, nil
}

func (s *AuthService) Login(ctx context.Context, login, password string) (string, error) {
	user, err := s.Users.FindByField(ctx, "login", login)
	if err != nil {
		return "", ErrInvalidCreds
	}

	if err = bcryptutil.CheckHash(password, user.PasswordHash); err != nil {
		return "", ErrInvalidCreds
	}

	return s.BuildJWTString(user.Login)
}

func (s *AuthService) ParseAndValidateToken(ctx context.Context, tokenString string) (*models.User, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(s.secretKey), nil
		})

	if err != nil || !token.Valid {
		return nil, ErrJWTToken
	}

	return s.Users.FindByField(ctx, "login", claims.UserLogin)
}

func (s *AuthService) Register(ctx context.Context, user *models.User) (string, error) {
	_, err := s.Users.FindByField(ctx, "login", user.Login)
	if err == nil {
		return "", ErrUserExists
	}

	hashed, err := bcryptutil.Hash(user.Password)
	if err != nil {
		return "", err
	}

	user.PasswordHash = hashed

	if err := s.Users.Create(ctx, user); err != nil {
		return "", err
	}

	return s.BuildJWTString(user.Login)
}
