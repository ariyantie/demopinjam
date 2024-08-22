package usecase

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/sync/errgroup"
	"io"
	"kredit-service/internal/consts"
	models "kredit-service/internal/model"
	"kredit-service/internal/repository"
	"kredit-service/internal/utils"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type UserHandler struct {
	u repository.UserRepository
	t repository.TransactionRepository
}

const (
	secret = "abc&1*~#^2^#s0^=)^^7%b34"
)

func NewUserUsecase(u repository.UserRepository, t repository.TransactionRepository) UserUcase {
	return &UserHandler{u, t}
}

func (u UserHandler) BulkApproveLoanRequest(ctx echo.Context, ids []int) (res []int, err error) {
	var (
		success []int
	)
	datas, err := u.u.CustomerLoanRequestByIds(ids, consts.LoanRequestStatusRequested)
	if err != nil {
		log.Errorf("[usecase][BulkApproveLoanRequest] error when CustomerLoanRequestByIds: %s", err.Error())
		return nil, err
	}
	for _, d := range datas {
		now := time.Now()
		expired := now.AddDate(0, d.Tenor, 0)
		err = u.u.UpdateLoanRequest(models.CustomerLoan{
			Status:       consts.LoanRequestStatusApproved,
			ApprovedDate: &now,
			ExpiredAt:    &expired,
			ID:           d.ID,
		})
		if err != nil {
			return nil, err
		}
		success = append(success, d.ID)
	}
	return success, nil
}

func (u UserHandler) ListRequestLoan(ctx echo.Context) ([]models.CustomerLoan, error) {
	return u.u.CustomerLoanRequest(consts.LoanRequestStatusRequested)
}

func (u UserHandler) UploadIdentity(ktp *multipart.FileHeader, selfie *multipart.FileHeader, userID int) error {
	var (
		err           error
		g             errgroup.Group
		uploadDir     = "./storage"
		encryptKTP    string
		encryptSelfie string
	)

	g.Go(func() error {
		src, err := ktp.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, os.ModePerm)
		}

		dstPath := filepath.Join(uploadDir, ktp.Filename)

		dst, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		encryptKTP, err = utils.Encrypt(ktp.Filename, secret)
		if err != nil {
			log.Errorf("error when encrypt ktp ")
			return err
		}
		return err
	})

	g.Go(func() error {
		src, err := selfie.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, os.ModePerm)
		}

		dstPath := filepath.Join(uploadDir, selfie.Filename)

		dst, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		encryptSelfie, err = utils.Encrypt(selfie.Filename, secret)
		if err != nil {
			log.Errorf("error when encrypt selfie ")
			return err
		}
		return err
	})

	if err = g.Wait(); err != nil {
		return err
	}

	err = u.u.UpdateIdentityUser(userID, encryptKTP, encryptSelfie)
	if err != nil {
		return err
	}

	return nil
}

func (u UserHandler) RequestLoan(ctx echo.Context, loan models.LoanRequestParam) error {
	now := time.Now()
	tenor, err := u.t.GetTenorByID(loan.TenorID)
	if err != nil {
		log.Errorf("[usecase][RequestLoan] error when GetTenorByID: %s", err.Error())
		return err
	}

	userInfo, err := u.u.GetUserByEmail(loan.Email)
	if err != nil {
		log.Errorf("[usecase][RequestLoan] error when GetUserByEmail: %s", err.Error())
		return err
	}

	err = u.u.RequestLoan(models.CustomerLoan{
		CustomerID: loan.CustomerID,
		Tenor:      loan.TenorID,
		LoanDate:   now,
		LoanAmount: tenor.Value * userInfo.Salary,
		Status:     consts.LoanRequestStatusRequested,
	})
	if err != nil {
		log.Errorf("[usecase][RequestLoan] error when RequestLoan: %s", err.Error())
		return err
	}
	return err
}

func (u UserHandler) GetUserLimit(ctx echo.Context, userID int) (models.LimitInformation, error) {
	limit, err := u.u.GetUserLimit(userID)
	if err != nil {
		log.Errorf("[usecase][GetUserLimit] error when GetUserLimit: %s", err.Error())
		return models.LimitInformation{}, err
	}
	if limit.ID == 0 {
		return models.LimitInformation{}, fmt.Errorf("you don't have a loan limit yet. Please apply for a loan")
	}
	return LoanToLimitInfo(limit), nil
}

func (u UserHandler) GetUserInfoByEmail(ctx echo.Context, email string) (models.Customer, error) {
	var (
		err error
		g   errgroup.Group
	)
	userInfo, err := u.u.GetUserByEmail(email)
	if err != nil {
		return models.Customer{}, err
	}
	g.Go(func() error {
		userInfo.NIK, err = utils.Decrypt(userInfo.NIK, secret)
		if err != nil {
			log.Errorf("error when decrypt nik ")
			return err
		}
		return err
	})

	g.Go(func() error {
		userInfo.FotoKTP, err = utils.Decrypt(userInfo.FotoKTP, secret)
		if err != nil {
			log.Errorf("error when decrypt KTP ")
			return err
		}
		return err
	})
	g.Go(func() error {
		userInfo.FotoSelfie, err = utils.Decrypt(userInfo.FotoSelfie, secret)
		if err != nil {
			log.Errorf("error when decrypt Selfie ")
			return err
		}
		return err
	})
	if err = g.Wait(); err != nil {
		return userInfo, err
	}
	return userInfo, err
}

func (u UserHandler) RegisterCustomer(ctx echo.Context, c models.CustomerParam) error {
	var (
		err error
		g   errgroup.Group
	)

	g.Go(func() error {
		// hash password
		c.Password, err = utils.HashPassword(c.Password)
		if err != nil {
			log.Errorf("error when hash password ")
			return err
		}
		return err
	})

	g.Go(func() error {
		//encrypt sensitive data
		c.NIK, err = utils.Encrypt(c.NIK, secret)
		if err != nil {
			log.Errorf("error when encrypt nik ")
			return err
		}
		return err
	})

	if err = g.Wait(); err != nil {
		return err
	}

	err = u.u.RegisterUser(c)
	if err != nil {
		log.Errorf("[usecase][RegisterCustomer] error when RegisterUser: %s", err.Error())
		return err
	}

	return nil
}
