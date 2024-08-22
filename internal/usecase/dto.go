package usecase

import (
	"kredit-service/internal/consts"
	models "kredit-service/internal/model"
)

func LoanToLimitInfo(cl models.CustomerLoan) models.LimitInformation {
	var res models.LimitInformation
	if cl.Status == consts.LoanRequestStatusRequested {
		res.Info = "loan request waiting approval"
		res.IsActive = false
	} else if cl.Status == consts.LoanRequestStatusApproved {
		res.Info = "your loan request already approved, you can use for transaction"
		res.IsActive = true
	} else if cl.Status == consts.LoanRequestStatusExpired {
		res.Info = "your loan request already expired, please contact admin to activate"
		res.IsActive = false
	} else {
		res.Info = "your limit already used for transaction"
		res.IsActive = true
	}
	res.UsedAmount = cl.UsedAmount
	res.AvailableAmount = cl.LoanAmount - cl.UsedAmount
	res.LoanAmount = cl.LoanAmount
	res.ExpiredDate = cl.ExpiredAt
	return res
}
