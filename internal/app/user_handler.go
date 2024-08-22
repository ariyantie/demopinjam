package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	models "kredit-service/internal/model"
	"kredit-service/internal/utils"
	"net/http"
	"time"
)

func (u handler) RegisterUser(c echo.Context) error {
	customer := new(models.CustomerParam)
	if err := c.Bind(customer); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}

	validator := validator.New()

	// Validasi struktur data customer
	if err := validator.Struct(customer); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "invalid payload",
			Error:    err.Error()})
	}

	err := u.User.RegisterCustomer(c, *customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, ResponseSuccess{
		Messages: "success register user",
	})
}

func (u handler) UploadKTPandSelfie(c echo.Context) error {
	// Validasi jenis file
	validImageTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
	}
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

	ktp, err := c.FormFile("ktp")
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "Gagal mendapatkan file ktp",
			Error:    "Foto ktp gagal dibaca",
		})
	}

	ktpType := ktp.Header.Get("Content-Type")
	if !validImageTypes[ktpType] {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "Gagal mengunggah file ktp",
			Error:    "Jenis file yang diunggah tidak diizinkan. Hanya file JPG atau PNG yang diizinkan.",
		})
	}

	selfie, err := c.FormFile("selfie")
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "Gagal mendapatkan file selfie",
			Error:    "Foto Selfie gagal dibaca",
		})
	}

	selfieType := selfie.Header.Get("Content-Type")
	if !validImageTypes[selfieType] {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "Gagal mengunggah file selfie",
			Error:    "Jenis file yang diunggah tidak diizinkan. Hanya file JPG atau PNG yang diizinkan.",
		})
	}
	err = u.User.UploadIdentity(ktp, selfie, int(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "gagal melakukan perubahan",
			Error:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success upload image",
	})
}

func (u handler) LoginUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	userInfo, err := u.User.GetUserInfoByEmail(c, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to login",
			Error:    err.Error(),
		})
	}

	if !utils.VerifyPassword(password, userInfo.Password) {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "invalid username/password",
			Error:    "username or password is mismatch",
		})
	}
	claims := &jwtCustomClaims{
		userInfo.Email,
		userInfo.IsAdmin,
		int64(userInfo.ID),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "error when generate token",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   accessToken,
	})
}

func (u handler) UserLimit(c echo.Context) error {
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
	data, err := u.User.GetUserLimit(c, int(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get user limit",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch user limit",
		Data:     data,
	})
}

func (u handler) GetCostumerProfile(c echo.Context) error {
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
	data, err := u.User.GetUserInfoByEmail(c, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get profile",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch user profile",
		Data:     data,
	})
}

func (u handler) RequestLoan(c echo.Context) error {
	loan := new(models.LoanRequestParam)
	if err := c.Bind(loan); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}

	validator := validator.New()
	// Validasi struktur data customer
	if err := validator.Struct(loan); err != nil {
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

	email, ok := claims["email"].(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get user email",
		})
	}
	loan.CustomerID = int(userID)
	loan.Email = email

	err := u.User.RequestLoan(c, *loan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to request loan",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, ResponseSuccess{
		Messages: "success request loan, please wait for approval",
	})
}
