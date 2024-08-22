package repository

import (
	"context"
	"database/sql"
	"fmt"
	models "kredit-service/internal/model"
	"strings"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserHandler{db}
}

func (h UserHandler) UpdateLoanRequest(loan models.CustomerLoan) error {
	_, err := h.db.Exec(queryApproveLoanRequest, loan.Status, loan.ApprovedDate, loan.ExpiredAt, loan.ID)
	if err != nil {
		return err
	}
	return err
}

func (h UserHandler) UpdateLoanRequestTx(tx *sql.Tx, loan models.CustomerLoan) error {
	_, err := tx.Exec(queryUseLimitRequest, loan.Status, loan.UsedAmount, loan.ID)
	if err != nil {
		return err
	}
	return err
}

func (h UserHandler) CustomerLoanRequestByIds(ids []int, status string) ([]models.CustomerLoan, error) {
	var (
		datas []models.CustomerLoan
		err   error
	)

	if len(ids) == 0 {
		return datas, nil
	}

	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "?"
	}
	inClause := strings.Join(placeholders, ",")

	query := fmt.Sprintf(queryGetListCustomerRequestByIds, inClause)

	args := make([]interface{}, len(ids)+1)
	args[0] = status
	for i, id := range ids {
		args[i+1] = id
	}

	rows, err := h.db.Query(query, args...)
	if err != nil {
		return datas, err
	}
	defer rows.Close()

	for rows.Next() {
		var data models.CustomerLoan
		if err = rows.Scan(&data.ID, &data.CustomerID, &data.Status, &data.UsedAmount, &data.LoanAmount, &data.Tenor, &data.CreatedAt, &data.UpdatedAt, &data.LoanDate); err != nil {
			return datas, err
		}
		datas = append(datas, data)
	}

	if err = rows.Err(); err != nil {
		return datas, err
	}
	return datas, err
}

func (h UserHandler) BeginTx() (*sql.Tx, error) {
	return h.db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
}

func (h UserHandler) CustomerLoanRequest(status string) ([]models.CustomerLoan, error) {
	var (
		datas []models.CustomerLoan
		err   error
	)
	rows, err := h.db.Query(queryGetListCustomerRequest, status)
	if err != nil {
		return datas, err
	}
	defer rows.Close()

	for rows.Next() {
		var data models.CustomerLoan
		if err = rows.Scan(&data.ID, &data.CustomerID, &data.Status, &data.UsedAmount, &data.LoanAmount, &data.Tenor, &data.CreatedAt, &data.UpdatedAt, &data.LoanDate); err != nil {
			return datas, err
		}
		datas = append(datas, data)
	}

	if err = rows.Err(); err != nil {
		return datas, err
	}
	return datas, err
}

func (h UserHandler) GetUserLimit(userID int) (models.CustomerLoan, error) {
	var (
		data models.CustomerLoan
		err  error
	)
	rows, err := h.db.Query(queryGetUserLimit, userID)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&data.ID, &data.CustomerID, &data.Status, &data.UsedAmount, &data.LoanAmount, &data.Tenor, &data.ExpiredAt); err != nil {
			return data, err
		}
	}

	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}

func (h UserHandler) RequestLoan(cl models.CustomerLoan) error {
	_, err := h.db.Exec(queryRequestLoan, cl.CustomerID, cl.Tenor, cl.LoanDate, cl.LoanAmount, cl.Status)
	if err != nil {
		return err
	}
	return err
}

func (h UserHandler) UpdateIdentityUser(id int, ktp string, selfie string) error {
	_, err := h.db.Exec(queryUpdateIdentity, ktp, selfie, id)
	if err != nil {
		return err
	}
	return err
}

func (h UserHandler) RegisterUser(c models.CustomerParam) error {
	_, err := h.db.Exec(insertNewCostumer, c.NIK, c.FullName, c.LegalName, c.BornPlace, c.BornDate, c.Salary, false, c.Email, c.Password, c.FotoSelfie, c.FotoKTP)
	if err != nil {
		return err
	}
	return err
}

func (h UserHandler) GetUserByEmail(email string) (models.Customer, error) {
	var (
		data models.Customer
		err  error
	)
	rows, err := h.db.Query(getCostumerByEmail, email)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&data.ID, &data.NIK, &data.FullName, &data.LegalName, &data.BornPlace, &data.BornDate, &data.Salary, &data.IsAdmin,
			&data.Email, &data.Password, &data.FotoSelfie, &data.FotoKTP,
			&data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
		); err != nil {
			return data, err
		}
	}

	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}
