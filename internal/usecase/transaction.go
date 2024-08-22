package usecase

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"kredit-service/internal/consts"
	models "kredit-service/internal/model"
	"kredit-service/internal/repository"
	"math"
	"time"
)

type TransactionHandler struct {
	t repository.TransactionRepository
	u repository.UserRepository
}

func NewTransactionsUsecase(t repository.TransactionRepository, u repository.UserRepository) TransactionUcase {
	return &TransactionHandler{t, u}
}

func (t TransactionHandler) GetTenorList(ctx echo.Context, email string) ([]models.Tenor, error) {
	tenors, err := t.t.GetTenorList()
	if err != nil {
		log.Errorf("[usecase][GetTenorList] error when get GetTenorList: %s", err.Error())
		return nil, err
	}

	user, err := t.u.GetUserByEmail(email)
	if err != nil {
		log.Errorf("[usecase][GetTenorList] error when GetUserByEmail: %s", err.Error())
		return nil, err
	}

	var res []models.Tenor

	for _, tenor := range tenors {
		amount := tenor.Value * user.Salary
		tenor.PossibleLimit = &amount
		res = append(res, tenor)
	}

	return res, err
}

func (t TransactionHandler) GetUserSchedulePayment(ctx echo.Context, userID int) ([]models.MonthPayments, error) {
	monthPayments, err := t.t.SchedulePayment(userID)
	if err != nil {
		log.Errorf("[usecase][GetUserSchedulePayment] error when get SchedulePayment: %s", err.Error())
		return nil, err
	}

	var res []models.MonthPayments

	for _, mp := range monthPayments {
		var total float64
		for _, p := range mp.Payments {
			total += p.Amount
		}

		res = append(res, models.MonthPayments{
			BillAmount: math.Round(total),
			Month:      mp.Month,
			Payments:   mp.Payments,
		})
	}
	return res, err
}

func (t TransactionHandler) PayTransaction(ctx echo.Context, param models.PaymentParam) error {
	_, err := time.Parse("2006-01", param.Date)
	if err != nil {
		log.Errorf("[usecase][PayTransaction] error when parsing param date: %s", err.Error())
		return fmt.Errorf("invalid date format, format should be YYYY-MM")
	}

	sch, err := t.t.SchedulePaymentByDate(param.UserID, param.Date)
	if err != nil {
		log.Errorf("[usecase][PayTransaction] error when SchedulePaymentByDate: %s", err.Error())
		return err
	}

	if len(sch) == 0 {
		return fmt.Errorf("you dont have bill on that month")
	}

	limit, err := t.u.GetUserLimit(param.UserID)
	if err != nil {
		log.Errorf("[usecase][PayTransaction] error when GetUserLimit: %s", err.Error())
		return err
	}

	var total float64
	var ids []int

	for _, s := range sch {
		total += s.Amount
		ids = append(ids, s.ID)
	}

	if math.Round(total) != param.Amount {
		return fmt.Errorf("amount payment should equal with bill amount")
	}

	tx, err := t.u.BeginTx()
	if err != nil {
		log.Errorf("[usecase][PayTransaction] error when BeginTx: %s", err.Error())
		return err
	}
	for _, id := range ids {
		err = t.t.UpdateSchedulePaymentTx(tx, models.SchedulePayment{
			Status:      consts.ScheduleStatusPaid,
			PaymentDate: time.Now(),
			ID:          id,
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = t.u.UpdateLoanRequestTx(tx, models.CustomerLoan{
		ID:         limit.ID,
		UsedAmount: limit.UsedAmount - param.Amount,
		Status:     consts.LoanRequestStatusUsed,
	})
	if err != nil {
		log.Errorf("[usecase][PayTransaction] error when UpdateLoanRequestTx: %s", err.Error())
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Errorf("[usecase][PayTransaction] error when Commit: %s", err.Error())
		return err
	}

	return nil
}

func (t TransactionHandler) CreateTransaction(ctx echo.Context, trx models.TransactionParam) error {
	limit, err := t.u.GetUserLimit(trx.UserID)
	if err != nil {
		log.Errorf("[usecase][CreateTransaction] error when GetUserLimit: %s", err.Error())
		return err
	}

	if limit.ID == 0 {
		return fmt.Errorf("you don't have a loan limit yet. Please apply for a loan")
	}

	trx.GenerateContractNumber()

	switch limit.Status {
	case consts.LoanRequestStatusExpired:
		err = fmt.Errorf("your limit already expired, please contact admin")
	case consts.LoanRequestStatusRequested:
		err = fmt.Errorf("your limit need approval, please wait before doing transaction")
	}
	if err != nil {
		return err
	}

	amount := limit.LoanAmount - limit.UsedAmount

	if trx.OTR > amount {
		return fmt.Errorf("available amount not enough to doing transaction")
	}

	amountWithInterest := trx.OTR + trx.AdminFee + (trx.Interest * trx.OTR)
	amountInstallment := amountWithInterest / float64(trx.TotalInstallment)
	trx.Status = consts.TransactionStatusSuccess

	//start operation using database transaction
	tx, err := t.u.BeginTx()
	if err != nil {
		return err
	}

	err = t.u.UpdateLoanRequestTx(tx, models.CustomerLoan{
		ID:         limit.ID,
		UsedAmount: limit.UsedAmount + amountWithInterest,
		Status:     consts.LoanRequestStatusUsed,
	})
	if err != nil {
		log.Errorf("[usecase][CreateTransaction] error when UpdateLoanRequestTx: %s", err.Error())
		tx.Rollback()
		return err
	}

	id, err := t.t.CreateTransactionTx(tx, trx)
	if err != nil {
		log.Errorf("[usecase][CreateTransaction] error when CreateTransactionTx: %s", err.Error())
		tx.Rollback()
		return err
	}

	for i := 1; i <= trx.TotalInstallment; i++ {
		now := time.Now()
		err = t.t.CreateSchedulePaymentTx(tx, models.SchedulePayment{
			TransactionID: id,
			Status:        consts.ScheduleStatusOnGoing,
			DueDate:       now.AddDate(0, i, 0),
			Amount:        amountInstallment,
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Errorf("[usecase][CreateTransaction] error when Commit: %s", err.Error())
		return err
	}
	return err

}
