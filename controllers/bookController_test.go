package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restful-api-testing/database"
	"restful-api-testing/models"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooksController(t *testing.T) {
	var testCases = []struct {
		name                    string
		path                    string
		expectedStatus          int
		expectedBodyStartWith   string
		expectedBodyMessageWith string
	}{{
		name:                    "succes get books",
		path:                    "/books",
		expectedStatus:          http.StatusOK,
		expectedBodyStartWith:   "{\"books\":[",
		expectedBodyMessageWith: "\"message\":\"success get all books\"}",
	},
	}

	e := InitEchoTestAPI()

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartWith))
			assert.True(t, strings.Contains(body, testCase.expectedBodyMessageWith))
		}
	}
}

func TestGetBookController(t *testing.T) {
	var testCases = []struct {
		name                    string
		path                    string
		expectedStatus          int
		expectedBodyStartWith   string
		expectedBodyMessageWith string
	}{{
		name:                    "succes get books",
		path:                    "/books",
		expectedStatus:          http.StatusOK,
		expectedBodyStartWith:   "{\"book\":[",
		expectedBodyMessageWith: "\"message\":\"success get book\"}",
	},
	}

	e := InitEchoTestAPI()

	book := database.SeedBook()
	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookID)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartWith))
			assert.True(t, strings.Contains(body, testCase.expectedBodyMessageWith))
		}
	}
}

func TestCreateBookController(t *testing.T) {
	var testCases = []struct {
		name                    string
		path                    string
		expectedStatus          int
		expectedBodyStartWith   string
		expectedBodyMessageWith string
	}{{
		name:                    "succes create book",
		path:                    "/books",
		expectedStatus:          http.StatusOK,
		expectedBodyStartWith:   "{\"book\":",
		expectedBodyMessageWith: "\"message\":\"success create new book\"}",
	},
	}

	e := InitEchoTestAPI()
	// token := database.SeedLogin()
	// bearer := "Bearer " + token
	book := database.SeedBook()
	jsonBody, _ := json.Marshal(&book)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/books", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", bearer)

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartWith))
			assert.True(t, strings.Contains(body, testCase.expectedBodyMessageWith))
		}
	}
}

func TestDeleteBookController(t *testing.T) {
	var testCases = []struct {
		name                    string
		path                    string
		expectedStatus          int
		expectedBodyStartWith   string
		expectedBodyMessageWith string
	}{{
		name:                    "succes create book",
		path:                    "/books",
		expectedStatus:          http.StatusOK,
		expectedBodyStartWith:   "{\"book\":",
		expectedBodyMessageWith: "\"message\":\"success delete book\"}",
	},
	}

	e := InitEchoTestAPI()
	// token := database.SeedLogin()
	// bearer := "Bearer " + token

	book := database.SeedBook()
	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodPut, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookID)
		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartWith))
			assert.True(t, strings.Contains(body, testCase.expectedBodyMessageWith))
		}
	}
}

func TestUpdateBookController(t *testing.T) {
	var testCases = []struct {
		name                    string
		path                    string
		expectedStatus          int
		expectedBodyStartWith   string
		expectedBodyMessageWith string
	}{{
		name:                    "succes create book",
		path:                    "/books",
		expectedStatus:          http.StatusOK,
		expectedBodyStartWith:   "{\"book\":",
		expectedBodyMessageWith: "\"message\":\"success update book\"}",
	},
	}

	e := InitEchoTestAPI()
	// token := database.SeedLogin()
	// bearer := "Bearer " + token
	book := database.SeedBook()
	bookInput := models.BookInput{
		Title:            "Test Update",
		Author:           "Test Update",
		Publisher:        "Test Update",
		Publication_Year: "Test Update",
		ISBN:             "Test Update",
		NumberOfPage:     "Test Update",
		Language:         "Test Update",
	}

	jsonBody, _ := json.Marshal(&bookInput)
	bodyReader := bytes.NewReader(jsonBody)

	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodPut, "/books", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", bearer)

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookID)
		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartWith))
			assert.True(t, strings.Contains(body, testCase.expectedBodyMessageWith))
		}
	}
}
