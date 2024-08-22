package models

import (
	"github.com/google/uuid"
	"time"
)

// Transaction adalah model struct untuk tabel transaction
type Transaction struct {
	ID               int        `json:"id"`
	UserID           int        `json:"user_id"`
	ContractNumber   string     `json:"contract_number"`
	OTR              float64    `json:"OTR"`
	AdminFee         float64    `json:"admin_fee"`
	TotalInstallment int        `json:"total_installment"`
	Interest         float64    `json:"interest"`
	AssetName        string     `json:"asset_name"`
	Status           string     `json:"status"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

type TransactionParam struct {
	UserID           int     `json:"-" `
	ContractNumber   string  `json:"-" `
	OTR              float64 `json:"OTR" validate:"required"`
	AdminFee         float64 `json:"admin_fee" validate:"required"`
	TotalInstallment int     `json:"total_installment" validate:"required"`
	Interest         float64 `json:"interest" validate:"required"`
	AssetName        string  `json:"asset_name" validate:"required"`
	Status           string  `json:"-"`
}

func (t *TransactionParam) GenerateContractNumber() {
	u := uuid.New()
	t.ContractNumber = u.String()
}

type SchedulePayment struct {
	ID            int        `json:"id"`
	TransactionID int        `json:"transaction_id"`
	PaymentDate   time.Time  `json:"payment_date"`
	Amount        float64    `json:"amount"`
	Status        string     `json:"status"`
	DueDate       time.Time  `json:"due_date"`
	LateFee       float64    `json:"late_fee"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}

type Payment struct {
	ID            int     `json:"id"`
	TransactionID int     `json:"transaction_id"`
	PaymentDate   string  `json:"payment_date"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	DueDate       string  `json:"due_date"`
	LateFee       float64 `json:"late_fee"`
}

type MonthPayments struct {
	BillAmount float64   `json:"bill_amount"`
	Month      string    `json:"month"`
	Payments   []Payment `json:"payments"`
}

type PaymentParam struct {
	Date   string  `json:"date" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
	UserID int     `json:"-"`
}
