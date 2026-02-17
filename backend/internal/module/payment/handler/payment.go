package handler

import (
	"encoding/json"
	"net/http"

	paymentUC "github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
)

type PaymentHandler struct {
	uc paymentUC.PaymentUsecase
}

func NewPaymentHandler(uc paymentUC.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{uc: uc}
}

func (h *PaymentHandler) GetDashboardV1Payments(
	w http.ResponseWriter,
	r *http.Request,
	params openapigen.GetDashboardV1PaymentsParams,
) {
	payments, err := h.uc.GetPayments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}
