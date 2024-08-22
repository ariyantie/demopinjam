package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	models "kredit-service/internal/model"
	"net/http"
)

func (u handler) TenorList(c echo.Context) error {
	// get user id from token
	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid or missing claims",
		})
	}

	email, ok := claims["email"].(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get user id",
		})
	}
	data, err := u.Transaction.GetTenorList(c, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get tenor list",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch tenor list",
		Data:     data,
	})
}

func (u handler) SchedulePayment(c echo.Context) error {
	// get user id from token
	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid or missing claims",
		})
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get user id",
		})
	}
	data, err := u.Transaction.GetUserSchedulePayment(c, int(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get schedule payment",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch Schedule Payment",
		Data:     data,
	})
}

func (u handler) PayTransaction(c echo.Context) error {
	transaction := new(models.PaymentParam)
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}

	validator := validator.New()

	if err := validator.Struct(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "invalid payload",
			Error:    err.Error()})
	}

	// get user id from token
	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid or missing claims",
		})
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get user id",
		})
	}

	transaction.UserID = int(userID)

	err := u.Transaction.PayTransaction(c, *transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to pay transaction",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success pay transaction",
	})
}

func (u handler) CreateTransaction(c echo.Context) error {
	transaction := new(models.TransactionParam)
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}

	validator := validator.New()

	if err := validator.Struct(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "invalid payload",
			Error:    err.Error()})
	}

	// get user id from token
	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid or missing claims",
		})
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get user id",
		})
	}

	transaction.UserID = int(userID)

	err := u.Transaction.CreateTransaction(c, *transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to create transaction",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, ResponseSuccess{
		Messages: "success doing transaction",
	})
}
