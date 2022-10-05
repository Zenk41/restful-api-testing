package controllers

import (
	"net/http"
	"net/http/httptest"
	"restful-api-testing/database"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestGetBooksController(t *testing.T) {
	var testCases = []struct {
		name                  string
		path                  string
		expectedStatus        int
		expectedBodyStartWith string
		expectedBodyMessageWith   string
	}{{
		name:                  "succes get books",
		path:                  "/books",
		expectedStatus:        http.StatusOK,
		expectedBodyStartWith: "{\"books\":",
		expectedBodyMessageWith:   "\"message\":\"success get all books\"",
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
		name                  string
		path                  string
		expectedStatus        int
		expectedBodyStartWith string
		expectedBodyMessageWith   string
	}{{
		name:                  "succes get books",
		path:                  "/books",
		expectedStatus:        http.StatusOK,
		expectedBodyStartWith: "{\"book\":",
		expectedBodyMessageWith:   "\"message\":\"success get book\"",
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
	
}
