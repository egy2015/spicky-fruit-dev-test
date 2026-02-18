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

func (h *PaymentHandler) PutDashboardV1PaymentsId(
	w http.ResponseWriter,
	r *http.Request,
	id int,
) {
	var req struct {
		MerchantName string `json:"merchant_name"`
		Amount       int    `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.MerchantName == "" || req.Amount <= 0 {
		http.Error(w, "invalid merchant_name or amount", http.StatusBadRequest)
		return
	}

	if err := h.uc.UpdatePayment(id, req.MerchantName, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"id":            id,
		"merchant_name": req.MerchantName,
		"amount":        req.Amount,
	})
}
