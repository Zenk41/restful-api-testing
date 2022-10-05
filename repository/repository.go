package repository

import "restful-api-testing/models"

type UserRepository interface {
	GetUsers() []models.User
	GetUserById(id int) models.User
	CreateUser(input models.UserInput) models.User
	UpdateUser(id int, input models.UserInput) models.User
	DeleteUser(id int) models.User
}

type BookRepository interface {
	GetBooks() []models.Book
	GetBookById(id int) models.Book
	CreateBook(input models.BookInput) models.Book
	UpdateBook(id int, input models.BookInput) models.Book
	DeleteBook(id int) models.Book
}

type AuthRepository interface {
	Register(input models.UserInput) (*models.User, error)
	Login(input models.UserLogin) string
}
