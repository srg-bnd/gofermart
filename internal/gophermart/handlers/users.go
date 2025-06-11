package handlers

import (
	"net/http"
	"strconv"
	"ya41-56/internal/gophermart/models"
	"ya41-56/internal/gophermart/repositories"
	"ya41-56/internal/gophermart/services"
	"ya41-56/internal/shared/httputil"
	"ya41-56/internal/shared/logger"
	"ya41-56/internal/shared/response"
)

type UsersHandler struct {
	Users repositories.Repository[models.User]
	Auth  *services.AuthService
}

func NewUsersHandler(authService *services.AuthService) *UsersHandler {
	return &UsersHandler{
		Users: authService.Users,
		Auth:  authService,
	}
}

func (h *UsersHandler) List(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page <= 0 {
		page = 1
	}

	perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
	if perPage <= 0 {
		perPage = 10
	}

	users, hasNext, err := h.Users.FindAllPaginated(r.Context(), repositories.Pagination{
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]interface{}{
		"list":    users,
		"hasNext": hasNext,
	})
}

func (h *UsersHandler) GetByID(w http.ResponseWriter, r *http.Request) {

}

func (h *UsersHandler) CreateNew(w http.ResponseWriter, r *http.Request) {
	type registerRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		ClientID string `json:"ClientID"`
		Level    int    `json:"UserLevel"`
		Status   int    `json:"status"`
	}

	var payload registerRequest

	if err := httputil.ParseJson(r, &payload); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid JSON payload")
		logger.L().Info(err.Error())
		return
	}

	createdUser, err := h.Auth.Register(r.Context(), &models.User{
		Login:    payload.Email,
		Password: payload.Password,
		Status:   models.UserStatus(payload.Status),
	})

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, createdUser)
}
