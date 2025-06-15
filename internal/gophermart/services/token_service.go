package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
	secretKey string
}

type Claims struct {
	jwt.RegisteredClaims
	UserLogin string `json:"login"`
}

const DefaultTokenExp = time.Hour * 1

func NewTokenService(secretKey string) *TokenService {
	return &TokenService{
		secretKey: secretKey,
	}
}

func (s *TokenService) BuildJWTString(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(DefaultTokenExp)),
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
