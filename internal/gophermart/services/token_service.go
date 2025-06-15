package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
	secretKey     string
	tokenLifetime time.Duration
}

type Claims struct {
	jwt.RegisteredClaims
	UserLogin string `json:"login"`
}

func NewTokenService(secretKey string, tokenLifetime time.Duration) *TokenService {
	return &TokenService{
		secretKey:     secretKey,
		tokenLifetime: tokenLifetime,
	}
}

func (s *TokenService) BuildJWTString(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenLifetime)),
		},
		UserLogin: login,
	})

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", ErrJWTToken
	}

	return tokenString, nil
}

func (s *TokenService) ParseToken(claims *Claims, tokenString string) (*jwt.Token, error) {
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

	return token, nil
}
