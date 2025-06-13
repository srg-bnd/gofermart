package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"ya41-56/internal/gophermart/models"
	"ya41-56/internal/gophermart/services"
	"ya41-56/internal/shared/contextutil"
	"ya41-56/internal/shared/response"
)

type AuthHandler struct {
	Auth *services.AuthService
}

func NewAuthHandler(auth *services.AuthService) *AuthHandler {
	return &AuthHandler{
		Auth: auth,
	}
}

// GetMe
func (h *AuthHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	userID, ok := contextutil.GetUserID(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	user, err := h.Auth.Users.FindByID(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	response.JSON(w, http.StatusOK, map[string]interface{}{
		"login":  user.Login,
		"status": user.Status,
	})
}

// Login
type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	jwtToken, err := h.Auth.Login(r.Context(), req.Login, req.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCreds) {
			response.Error(w, http.StatusUnauthorized, err.Error())
		} else {
			response.Error(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	w.Header().Set("Authorization", "Bearer "+jwtToken)
	response.JSON(w, http.StatusOK, nil)
}

// Ping
func (h *AuthHandler) ProtectedPing(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("authorized"))
	if err != nil {
		return
	}
}

// Register
type registerRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	jwtToken, err := h.Auth.Register(r.Context(), &models.User{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, services.ErrUserExists) {
			response.Error(w, http.StatusConflict, err.Error())
		} else {
			response.Error(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	w.Header().Set("Authorization", "Bearer "+jwtToken)
	response.JSON(w, http.StatusOK, nil)
}
