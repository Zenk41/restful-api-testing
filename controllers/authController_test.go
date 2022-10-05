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
		name                     string
		path                     string
		expectCode               int
		expectedBodyStartsWith   string
		expectedBodyContainsWith string
	}{
		{
			name:                     "register normal",
			path:                     "/register",
			expectCode:               http.StatusCreated,
			expectedBodyStartsWith:   "{\"message\":\"success register new user\"",
			expectedBodyContainsWith: ",\"user\":[",
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
			assert.True(t, strings.Contains(body, testCase.expectedBodyContainsWith))
		}
	}
}

func TestUserLogin(t *testing.T) {

	var testCases = []struct {
		name                 string
		path                 string
		expectCode           int
		expectedBodyStartsWith string
	}{
		{
			name:                 "get token normal",
			path:                 "/login",
			expectCode:           http.StatusCreated,
			expectedBodyStartsWith: "{\"token\":",
		},
	}

	e := InitEchoTestAPI()
	userLogin := database.SeedRegister()
	token := database.SeedLogin()
	bearer := "Bearer " + token

	jsonBody, _ := json.Marshal(&userLogin)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/login", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, Login(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}
