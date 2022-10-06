package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restful-api-testing/database"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	database.InitDBTest()
	database.InitMigrateTest()
	e := echo.New()
	return e
}

func TestRegisterUser(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectCode             int
		expectedBodyStartsWith string
	}{
		{
			name:                   "register normal",
			path:                   "/register",
			expectCode:             http.StatusCreated,
			expectedBodyStartsWith: "{\"message\":\"success register new user\"",
		},
	}

	e := InitEchoTestAPI()

	userInput := database.SeedRegister()

	jsonBody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/register", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, Register(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestUserLogin(t *testing.T) {

	var testCases = []struct {
		name                   string
		path                   string
		expectCode             int
		expectedBodyStartsWith string
	}{
		{
			name:                   "get token normal",
			path:                   "/login",
			expectCode:             http.StatusOK,
			expectedBodyStartsWith: "{\"token\":",
		},
	}

	e := InitEchoTestAPI()

	database.SeedRegister()
	userLogin, _ := database.SeedLogin()

	jsonBody, _ := json.Marshal(&userLogin)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/login", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}
