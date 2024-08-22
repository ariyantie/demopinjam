package repository

import (
	"database/sql"
	"encoding/json"
	models "kredit-service/internal/model"
)

type TransactionHandler struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &TransactionHandler{db}
}

func (h TransactionHandler) CreateSchedulePaymentTx(tx *sql.Tx, data models.SchedulePayment) error {
	_, err := tx.Exec(queryCreateSchedulePayment, data.TransactionID, data.Amount, data.Status, data.DueDate)
	if err != nil {
		return err
	}
	return err
}

func (h TransactionHandler) SchedulePayment(userID int) ([]models.MonthPayments, error) {
	var (
		data []models.MonthPayments
		err  error
	)

	rows, err := h.db.Query(queryGetListTransaction, userID)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var month string
		var paymentsJSON string
		if err = rows.Scan(&month, &paymentsJSON); err != nil {
			return data, err
		}

		var payments []models.Payment
		if err = json.Unmarshal([]byte(paymentsJSON), &payments); err != nil {
			return data, err
		}

		data = append(data, models.MonthPayments{
			Month:    month,
			Payments: payments,
		})
	}

	return data, err
}

func (h TransactionHandler) CreateTransactionTx(tx *sql.Tx, data models.TransactionParam) (int, error) {
	result, err := tx.Exec(queryCreateTransaction, data.UserID, data.ContractNumber, data.OTR, data.AdminFee, data.TotalInstallment, data.Interest, data.AssetName, data.Status)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), err
}

func (h TransactionHandler) UpdateSchedulePaymentTx(tx *sql.Tx, data models.SchedulePayment) error {
	_, err := tx.Exec(queryUpdateSchedulePayment, data.Status, data.PaymentDate, data.ID)
	if err != nil {
		return err
	}

	return err
}

func (h TransactionHandler) GetTenorByID(id int) (models.Tenor, error) {
	var (
		data models.Tenor
		err  error
	)
	rows, err := h.db.Query(getTenorByID, id)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&data.ID, &data.Tenor, &data.Value); err != nil {
			return data, err
		}
	}

	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}

func (h TransactionHandler) SchedulePaymentByDate(userID int, date string) ([]models.SchedulePayment, error) {
	var (
		data []models.SchedulePayment
		err  error
	)
	rows, err := h.db.Query(queryGetSchedulePayment, userID, date)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var schedule models.SchedulePayment
		if err = rows.Scan(&schedule.ID, &schedule.TransactionID, &schedule.PaymentDate, &schedule.Amount, &schedule.Status,
			&schedule.DueDate, &schedule.LateFee, &schedule.CreatedAt, &schedule.UpdatedAt, &schedule.DeletedAt); err != nil {
			return data, err
		}
		data = append(data, schedule)
	}

	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}

func (h TransactionHandler) GetUserTransactionByUserID(userID int) ([]models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (h TransactionHandler) GetTenorList() ([]models.Tenor, error) {
	var (
		data []models.Tenor
		err  error
	)
	rows, err := h.db.Query(getTenorList)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Tenor
		if err = rows.Scan(&t.ID, &t.Tenor, &t.Value); err != nil {
			return data, err
		}
		data = append(data, t)
	}

	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}
