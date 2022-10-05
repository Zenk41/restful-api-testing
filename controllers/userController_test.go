package controllers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"restful-api-testing/database"
// 	"restful-api-testing/models"
// 	"strings"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func InsertDataUserForGetUsers() error {
// 	user := models.UserInput{
// 		Name:     "Alta",
// 		Password: "123",
// 		Email:    "alta@gmail.com",
// 	}

// 	var err error
// 	if err = database.DB.Save(&user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func TestCreateUserController(t *testing.T) {
// 	var testCases = []struct {
// 		name                 string
// 		path                 string
// 		expectCode           int
// 		expectBodyStartsWith string
// 	}{
// 		{
// 			name:                 "get users normal",
// 			path:                 "/users/:id",
// 			expectCode:           http.StatusOK,
// 			expectBodyStartsWith: "{\"message\":\"success creating new user\",\"user\":[",
// 		},
// 	}
// 	e := InitEchoTestAPI()
// 	InsertDataUserForGetUsers()
// 	req := httptest.NewRequest(http.MethodPost, "/", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)

// 		if assert.NoError(t, GetUsersController(c)) {
// 			assert.Equal(t, testCase.expectCode, rec.Code)
// 			body := rec.Body.String()

// 			// open file
// 			// convert struct
// 			var user UserResponse
// 			err := json.Unmarshal([]byte(body), &user)
// 			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))

// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 		}
// 	}
// }

// func TestGetUsersController(t *testing.T) {
// 	var testCases = []struct {
// 		name                 string
// 		path                 string
// 		expectCode           int
// 		expectBodyStartsWith string
// 	}{
// 		{
// 			name:                 "get users normal",
// 			path:                 "/users",
// 			expectCode:           http.StatusOK,
// 			expectBodyStartsWith: "{\"message\":\"success get all users\",\"users\":[",
// 		},
// 	}

// 	e := InitEchoTestAPI()
// 	InsertDataUserForGetUsers()
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)

// 		if assert.NoError(t, GetUsersController(c)) {
// 			assert.Equal(t, testCase.expectCode, rec.Code)
// 			body := rec.Body.String()

// 			// open file
// 			// convert struct
// 			var user UserResponse
// 			err := json.Unmarshal([]byte(body), &user)
// 			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))

// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 		}
// 	}
// }

// func TestGetUserController(t *testing.T) {
// 	var testCases = []struct {
// 		name                 string
// 		path                 string
// 		expectCode           int
// 		expectBodyStartsWith string
// 	}{
// 		{
// 			name:                 "get user normal",
// 			path:                 "/users/:id",
// 			expectCode:           http.StatusOK,
// 			expectBodyStartsWith: "{\"message\":\"success get user\",\"user\":[",
// 		},
// 	}
	
// 	e := InitEchoTestAPI()
// 	InsertDataUserForGetUsers()
// 	req := httptest.NewRequest(http.MethodGet, "/:id", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)

// 		if assert.NoError(t, GetUserController(c)) {
// 			assert.Equal(t, testCase.expectCode, rec.Code)
// 			body := rec.Body.String()

// 			// open file
// 			// convert struct
// 			var user UserResponse
// 			err := json.Unmarshal([]byte(body), &user)
// 			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))

// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 		}
// 	}
// }

// func TestDeleteUserController(t *testing.T) {

// }

// func TestUpdateUserController(t *testing.T) {

// }
