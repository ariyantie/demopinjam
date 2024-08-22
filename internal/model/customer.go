package models

import (
	"time"
)

// CustomerLoan adalah model struct untuk tabel customer_loan
type CustomerLoan struct {
	ID           int        `json:"id"`
	CustomerID   int        `json:"customer_id"`
	Status       string     `json:"status"`
	LoanDate     time.Time  `json:"loan_date,omitempty"`
	LoanAmount   float64    `json:"loan_amount"`
	UsedAmount   float64    `json:"used_amount"`
	ApprovedDate *time.Time `json:"approved_date,omitempty"`
	Tenor        int        `json:"tenor"`
	ExpiredAt    *time.Time `json:"expired_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

type LoanRequestParam struct {
	Email      string
	CustomerID int
	TenorID    int `json:"tenor_id"`
}

type LimitInformation struct {
	AvailableAmount float64    `json:"available_amount"`
	UsedAmount      float64    `json:"used_amount"`
	LoanAmount      float64    `json:"loan_amount"`
	Info            string     `json:"info"`
	IsActive        bool       `json:"is_active"`
	ExpiredDate     *time.Time `json:"expired_date,omitempty"`
}

// Customer adalah model struct untuk tabel customer
type Customer struct {
	ID         int        `json:"id" db:"id"`
	NIK        string     `json:"NIK" db:"NIK"`
	FullName   string     `json:"full_name" db:"full_name"`
	LegalName  string     `json:"legal_name,omitempty" db:"legal_name"`
	BornPlace  string     `json:"born_place,omitempty" db:"born_place"`
	BornDate   time.Time  `json:"born_date,omitempty" db:"born_date"`
	Salary     float64    `json:"salary,omitempty" db:"salary"`
	Email      string     `json:"email" db:"email"`
	IsAdmin    bool       `json:"is_admin,omitempty" db:"is_admin"`
	Password   string     `json:"password" db:"password"`
	FotoSelfie string     `json:"foto_selfie,omitempty" db:"foto_selfie"`
	FotoKTP    string     `json:"foto_ktp,omitempty" db:"foto_ktp"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CustomerParam adalah model struct untuk request
type CustomerParam struct {
	NIK        string    `json:"NIK" validate:"required"`
	FullName   string    `json:"full_name" validate:"required"`
	LegalName  string    `json:"legal_name,omitempty" validate:"required"`
	BornPlace  string    `json:"born_place,omitempty" validate:"required"`
	BornDate   time.Time `json:"born_date,omitempty" validate:"required"`
	Salary     float64   `json:"salary,omitempty" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required"`
	FotoSelfie string    `json:"-"`
	FotoKTP    string    `json:"-"`
}

// LoanPayment adalah model struct untuk tabel loan_payment
type LoanPayment struct {
	ID             int        `json:"id"`
	CustomerLoanID int        `json:"customer_loan_id"`
	PaymentDate    time.Time  `json:"payment_date"`
	Amount         float64    `json:"amount"`
	Status         string     `json:"status"`
	DueDate        time.Time  `json:"due_date"`
	LateFee        float64    `json:"late_fee"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}

type CustomerLimit struct {
}
