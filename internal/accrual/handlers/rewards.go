package handlers

import (
	"encoding/json"
	"net/http"
	"ya41-56/internal/accrual/models"
	"ya41-56/internal/shared/repositories"
	"ya41-56/internal/shared/response"
)

type RewardsHandler struct {
	Repo repositories.Repository[models.RewardMechanic]
}

func NewRewardsHandler(repo repositories.Repository[models.RewardMechanic]) *RewardsHandler {
	return &RewardsHandler{
		Repo: repo,
	}
}

type CreateMechanicRequest struct {
	Match      string  `json:"match"`
	Reward     float32 `json:"reward"`
	RewardType string  `json:"reward_type"`
}

func (h *RewardsHandler) CreateRewards(w http.ResponseWriter, r *http.Request) {
	var req CreateMechanicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if req.Match == "" || (req.RewardType != "pt" && req.RewardType != "%") {
		response.Error(w, http.StatusUnprocessableEntity, "invalid mechanic data")
		return
	}

	reward := models.RewardMechanic{
		Match:      req.Match,
		Reward:     req.Reward,
		RewardType: req.RewardType,
	}

	if err := h.Repo.Create(r.Context(), &reward); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]uint{
		"id": reward.ID,
	})
}
