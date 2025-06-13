package handlers

import (
	"net/http"
	"ya41-56/internal/shared/response"
)

type UsersHandler struct {
}

func NewUsersHandler() *UsersHandler {
	return &UsersHandler{}
}

func (h *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	response.Error(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func (h *UsersHandler) Register(w http.ResponseWriter, r *http.Request) {
	response.Error(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}
