package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"ya41-56/internal/gophermart/services"
	"ya41-56/internal/shared/contextutil"
	"ya41-56/internal/shared/response"
)

type AuthMiddleware struct {
	Auth *services.AuthService
}

func New(auth *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		Auth: auth,
	}
}

func (m *AuthMiddleware) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		currentUser, err := m.Auth.ParseAndValidate(r.Context(), token)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		ctx := r.Context()
		ctx = contextutil.WithUserID(ctx, strconv.Itoa(int(currentUser.ID)))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
