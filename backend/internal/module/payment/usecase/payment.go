package usecase

import (
	"database/sql"
)

type Payment struct {
	ID           int    `json:"id"`
	Amount       int    `json:"amount"`
	MerchantName string `json:"merchant_name"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
}

type PaymentUsecase interface {
	GetPayments() ([]Payment, error)
	UpdatePayment(id int, merchantName string, amount int) error
}

type paymentUsecase struct {
	db *sql.DB
}

func NewPaymentUsecase(db *sql.DB) PaymentUsecase {
	return &paymentUsecase{db: db}
}

func (u *paymentUsecase) GetPayments() ([]Payment, error) {
	rows, err := u.db.Query(`
		SELECT id, merchant_name, amount, status, created_at
		FROM payments
		ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []Payment

	for rows.Next() {
		var p Payment
		if err := rows.Scan(
			&p.ID,
			&p.MerchantName,
			&p.Amount,
			&p.Status,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}

	return payments, nil
}

func (u *paymentUsecase) UpdatePayment(
	id int,
	merchantName string,
	amount int,
) error {
	_, err := u.db.Exec(`
		UPDATE payments
		SET merchant_name = ?, amount = ?
		WHERE id = ?
	`, merchantName, amount, id)

	return err
}
