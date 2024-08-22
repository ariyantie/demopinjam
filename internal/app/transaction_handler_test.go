package app

import (
	models "kredit-service/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// MockTransactionUsecase is a mock implementation of the transaction usecase.
type MockTransactionUsecase struct {
}

func (m *MockTransactionUsecase) GetUserSchedulePayment(ctx echo.Context, userID int) ([]models.MonthPayments, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockTransactionUsecase) PayTransaction(ctx echo.Context, param models.PaymentParam) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockTransactionUsecase) CreateTransaction(ctx echo.Context, transaction models.TransactionParam) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockTransactionUsecase) GetTenorList(ctx echo.Context, email string) ([]models.Tenor, error) {
	// Simulate the behavior of the transaction usecase.
	// Return the appropriate data or error based on your test case.
	// For example:
	return []models.Tenor{}, nil
}

func TestTenorList(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	h := handler{Transaction: &MockTransactionUsecase{}}

	req := httptest.NewRequest(http.MethodGet, "/customer/tenor", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := h.TenorList(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
