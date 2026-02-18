package handler

import (
	"encoding/json"
	"net/http"

	userUC "github.com/durianpay/fullstack-boilerplate/internal/module/user/usecase"
)

type UserHandler struct {
	uc userUC.UserUsecase
}

func NewUserHandler(uc userUC.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) GetDashboardV1Users(w http.ResponseWriter, r *http.Request) {
	users, err := h.uc.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
