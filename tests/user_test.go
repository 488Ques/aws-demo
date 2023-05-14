package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/488Ques/aws-demo/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]*models.User{
		// "jon@labstack.com": {Name: "Jon Snow", Email: "jon@labstack.com"},
	}
	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
	e        = echo.New()
	h        = &models.UserModel{DB: mockDB}
)

func TestCreateUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, strings.TrimSpace(rec.Body.String()))
		assert.Equal(t, mockDB["jon@labstack.com"], &models.User{Name: "Jon Snow", Email: "jon@labstack.com"})
	}
}

func TestGetUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/:email")
	c.SetParamNames("email")
	c.SetParamValues("jon@labstack.com")

	if assert.NoError(t, h.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, strings.TrimSpace(rec.Body.String()))
	}
}
