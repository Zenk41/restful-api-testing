package services

import (
	"restful-api-testing/database"
	"restful-api-testing/models"
)

// get all books
func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := database.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// get book by id
func GetBookById(id int) (interface{}, error) {
	var book []models.Book
	if err := database.DB.First(&book, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// create new book
func CreateBook(book models.Book) (interface{}, error) {
	if err := database.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// update book by id
func UpdateBook(id int, book models.Book) (interface{}, error) {
	books := models.Book{}
	if err := database.DB.Model(&books).Where("id=?", id).Updates(&book).First(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// delete book by id
func DeleteBook(id int) (interface{}, error) {
	var book []models.Book
	if err := database.DB.Where("id=?", id).First(&book).Delete(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}
