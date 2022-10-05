package repository

import (
	"restful-api-testing/database"
	"restful-api-testing/models"
)

type BookRepo struct{}

// get all books
func (b *BookRepo) GetBooks() ([]models.Book, error) {
	var books []models.Book

	if err := database.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// get book by id
func (b *BookRepo) GetBookById(id int) (*models.Book, error) {
	var book models.Book
	if err := database.DB.First(&book, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

// create new book
func (b *BookRepo) CreateBook(input models.BookInput) (*models.Book, error) {
	var newBook models.Book = models.Book{
		Title:            input.Title,
		Author:           input.Author,
		Publisher:        input.Publisher,
		Publication_Year: input.Publication_Year,
		ISBN:             input.ISBN,
		NumberOfPage:     input.NumberOfPage,
		Language:         input.Language,
	}
	if err := database.DB.Create(&newBook).Error; err != nil {
		return nil, err
	}
	return &newBook, nil
}

// update book by id
func (b *BookRepo) UpdateBook(id int, input models.BookInput) (*models.Book, error) {
	book, err := b.GetBookById(id)
	if err != nil {
		return nil, err
	}
	book.Title = input.Title
	book.Author = input.Author
	book.Publisher = input.Publisher
	book.Publication_Year = input.Publication_Year
	book.ISBN = input.ISBN
	book.NumberOfPage = input.NumberOfPage
	book.Language = input.Language
	if err := database.DB.Save(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

// delete book by id
func (b *BookRepo) DeleteBook(id int) (*models.Book, error) {
	book, err := b.GetBookById(id)
	if err != nil {
		return nil, err
	}
	if err := database.DB.Delete(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}
