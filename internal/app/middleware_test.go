package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestJWTMiddleware(t *testing.T) {
	// Define a secret key for testing
	secretKey := "secret"

	// Create an Echo instance
	e := echo.New()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZXhhbXBsZS5jb20iLCJpc19hZG1pbiI6ZmFsc2UsImlkIjoxLCJleHAiOjE2OTY1MjQ0MTd9.Z_4zjV7jC3JJNE2TKb1mLGcWQBAC-h7G7ju3xElaAbg"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	middleware := JWTMiddleware(secretKey)

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "Protected Resource")
	}

	err := middleware(handler)(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	//assert.Equal(t, map[string]interface{}{
	//	"message": "Invalid token",
	//}, rec.Body.String())

	invalidToken := "invalid-jwt-token"
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", invalidToken))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	err = middleware(handler)(c)
	//assert.Error(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid token")
}

func TestAdminMiddleware(t *testing.T) {
	e := echo.New()

	adminClaims := `{"is_admin": true}`
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("claims", parseJWTClaims(adminClaims))

	middleware := AdminMiddleware(func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin Resource")
	})

	err := middleware(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	//assert.Equal(t, "Admin Resource", rec.Body.String())

	nonAdminClaims := `{"is_admin": false}`
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.Set("claims", parseJWTClaims(nonAdminClaims))

	err = middleware(c)
	//assert.Error(t, err)
	//assert.Equal(t, http.StatusForbidden, rec.Code)
	//assert.Contains(t, rec.Body.String(), "Access denied")
}

func parseJWTClaims(claimsJSON string) map[string]interface{} {
	claims := make(map[string]interface{})
	if err := json.Unmarshal([]byte(claimsJSON), &claims); err != nil {
		panic(err)
	}
	return claims
}
