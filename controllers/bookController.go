package controllers

import (
	"net/http"
	"restful-api-testing/models"
	"restful-api-testing/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	books, err := services.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := services.GetBookById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	books := models.Book{}
	if err := c.Bind(&books); err != nil {
		return err
	}
	book, err := services.CreateBook(books)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := services.DeleteBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
		"book":    book,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	books := models.Book{}
	if err := c.Bind(&books); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := services.UpdateBook(id, books)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}
