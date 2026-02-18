package api

import (
	"net/http"

	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	paymenthandler "github.com/durianpay/fullstack-boilerplate/internal/module/payment/handler"
	userhandler "github.com/durianpay/fullstack-boilerplate/internal/module/user/handler"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
)

type APIHandler struct {
	Auth    *ah.AuthHandler
	Payment *paymenthandler.PaymentHandler
	User    *userhandler.UserHandler
}

var _ openapigen.ServerInterface = (*APIHandler)(nil)

func (h *APIHandler) PostDashboardV1AuthLogin(w http.ResponseWriter, r *http.Request) {
	h.Auth.PostDashboardV1AuthLogin(w, r)
}

func (h *APIHandler) GetDashboardV1Payments(w http.ResponseWriter, r *http.Request, params openapigen.GetDashboardV1PaymentsParams) {
	h.Payment.GetDashboardV1Payments(w, r, params)
}

func (h *APIHandler) GetDashboardV1Users(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.User.GetDashboardV1Users(w, r)
}

func (h *APIHandler) PutDashboardV1PaymentsId(
	w http.ResponseWriter,
	r *http.Request,
	id int,
) {
	h.Payment.PutDashboardV1PaymentsId(w, r, id)
}
