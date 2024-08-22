package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func (u handler) ListCostumerLoan(c echo.Context) error {
	data, err := u.User.ListRequestLoan(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get list loan request",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch loan request list",
		Data:     data,
	})
}

func (u handler) BulkApproveLoanRequest(c echo.Context) error {
	idParam := c.QueryParam("id")
	idParam = strings.TrimSuffix(strings.TrimPrefix(idParam, "["), "]")
	idStrings := strings.Split(idParam, ",")
	var idInts []int
	for _, str := range idStrings {
		id, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		idInts = append(idInts, id)
	}
	data, err := u.User.BulkApproveLoanRequest(c, idInts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to  bulk approval",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success bulk approval",
		Data:     data,
	})
}
