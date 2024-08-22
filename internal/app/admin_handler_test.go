package app

import (
	"errors"
	models "kredit-service/internal/model"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserUsecase is a mock implementation of the user usecase.
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) UploadIdentity(ktp *multipart.FileHeader, selfie *multipart.FileHeader, user_id int) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserUsecase) RegisterCustomer(ctx echo.Context, customer models.CustomerParam) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserUsecase) GetUserInfoByEmail(ctx echo.Context, email string) (models.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserUsecase) GetUserLimit(ctx echo.Context, userID int) (models.LimitInformation, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserUsecase) RequestLoan(ctx echo.Context, loan models.LoanRequestParam) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserUsecase) BulkApproveLoanRequest(ctx echo.Context, ids []int) (res []int, err error) {
	if len(ids) == 0 {
		return nil, errors.New("empty ID list")
	}
	return []int{}, nil
}

func (m *MockUserUsecase) ListRequestLoan(c echo.Context) ([]models.CustomerLoan, error) {
	args := m.Called(c)
	return args.Get(0).([]models.CustomerLoan), args.Error(1)
}

func TestListCostumerLoan(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	mockUsecase := new(MockUserUsecase)

	h := handler{User: mockUsecase}

	expectedData := []models.CustomerLoan{
		models.CustomerLoan{
			ID: 1,
		},
	}

	mockUsecase.On("ListRequestLoan", mock.Anything).Return(expectedData, nil)

	req := httptest.NewRequest(http.MethodGet, "/list-loan", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := h.ListCostumerLoan(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertCalled(t, "ListRequestLoan", mock.Anything)
}

func TestBulkApproveLoanRequest(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()
	h := handler{User: &MockUserUsecase{}}

	queryParam := "[1,2,3]"

	req := httptest.NewRequest(http.MethodPut, "/admin/approve-loan?id="+queryParam, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := h.BulkApproveLoanRequest(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
