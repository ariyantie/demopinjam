package app

import (
	models "kredit-service/internal/model"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Define a mock UserUcase and TransactionUcase for testing
type mockUserUcase struct{}

func (u *mockUserUcase) UploadIdentity(ktp *multipart.FileHeader, selfie *multipart.FileHeader, user_id int) error {
	//TODO implement me
	panic("implement me")
}

func (u *mockUserUcase) GetUserLimit(ctx echo.Context, userID int) (models.LimitInformation, error) {
	//TODO implement me
	panic("implement me")
}

func (u *mockUserUcase) ListRequestLoan(ctx echo.Context) ([]models.CustomerLoan, error) {
	//TODO implement me
	panic("implement me")
}

func (u *mockUserUcase) BulkApproveLoanRequest(ctx echo.Context, ids []int) (res []int, err error) {
	//TODO implement me
	panic("implement me")
}

type mockTransactionUcase struct{}

func (u *mockTransactionUcase) GetTenorList(ctx echo.Context, email string) ([]models.Tenor, error) {
	//TODO implement me
	panic("implement me")
}

func (u *mockTransactionUcase) GetUserSchedulePayment(ctx echo.Context, userID int) ([]models.MonthPayments, error) {
	//TODO implement me
	panic("implement me")
}

func (u *mockTransactionUcase) PayTransaction(ctx echo.Context, param models.PaymentParam) error {
	//TODO implement me
	panic("implement me")
}

func (u *mockUserUcase) RegisterCustomer(c echo.Context, customer models.CustomerParam) error {
	return nil
}

func (u *mockUserUcase) GetUserInfoByEmail(c echo.Context, email string) (models.Customer, error) {
	return models.Customer{}, nil
}

func (u *mockUserUcase) RequestLoan(c echo.Context, loan models.LoanRequestParam) error {
	return nil
}

func (u *mockTransactionUcase) CreateTransaction(c echo.Context, transaction models.TransactionParam) error {
	return nil
}

func TestLoginUser(t *testing.T) {
	e := echo.New()
	payload := `{"email":"johndoe@example.com","password":"secure_password"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	u := handler{User: &mockUserUcase{}}

	err := u.LoginUser(c)
	assert.NoError(t, err)
	//assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUserLimit(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/customer/limit", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("claims", map[string]interface{}{"id": float64(1)})
	u := handler{User: &mockUserUcase{}}

	err := u.UserLimit(c)
	assert.NoError(t, err)
	//assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestRequestLoan(t *testing.T) {
	e := echo.New()
	payload := `{"loanAmount":1000.0,"tenor":12}`
	req := httptest.NewRequest(http.MethodPost, "/customer/request-loan", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("claims", map[string]interface{}{"id": float64(1), "email": "john@example.com"})
	u := handler{User: &mockUserUcase{}}

	err := u.RequestLoan(c)
	assert.NoError(t, err)
	//assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestCreateTransaction(t *testing.T) {
	e := echo.New()
	payload := `{"userID":1,"amount":1000.0}`
	req := httptest.NewRequest(http.MethodPost, "/customer/create-transaction", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("claims", map[string]interface{}{"id": float64(1)})
	u := handler{Transaction: &mockTransactionUcase{}}

	err := u.CreateTransaction(c)
	assert.NoError(t, err)
	//assert.Equal(t, http.StatusCreated, rec.Code)
}
