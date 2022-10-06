package database

import (
	"errors"
	"fmt"
	"log"
	mid "restful-api-testing/middlewares"
	"restful-api-testing/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// database configuration
type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

// initialize database
func InitDB() {

	config := Config{
		DB_Username: "root",
		DB_Password: "@root123",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// auto migrate tabel using struct in model
func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}

// initialize database for testing
func InitDBTest() {
	config := Config{
		DB_Username: "root",
		DB_Password: "@root123",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go_test",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// auto migrate tabel using struct in model and deleting existing table
func InitMigrateTest() {
	DB.Migrator().DropTable(&models.User{}, &models.Book{})
	DB.AutoMigrate(&models.User{}, &models.Book{})
}


// Seed
// Seed for register
func SeedRegister() models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)

	var newUser models.User = models.User{
		Name:     "Ardhi Ramadhani",
		Email:    "ardhidhani@gmail.com",
		Password: string(password),
	}
	if err := DB.Create(&newUser).Error; err != nil {
		panic(err)
	}
	var createdUser models.User
	DB.Last(&createdUser)
	return createdUser
}

// Seed for Login to get token
func SeedLogin() (*models.User, string) {
	user := SeedRegister()
	var loginUser models.UserLogin= models.UserLogin{
	Email: "ardhidhani@gmail.com",
	Password: "123",
	}
	if err := DB.First(&user, "email=?", loginUser.Email).Error; err != nil {
		panic(err)
	}
	if user.ID == 0 {
		return nil,""
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))

	if err != nil {
		return nil,""
	}
	token, err := mid.CreateToken(user.ID)
	if err != nil {
		return nil,""
	}
	return &user, token
}

// Seed for Creating Book
func SeedBook() models.Book {
	var bookInput models.Book = models.Book{
		Title:            "Atomic Habits: An Easy & Proven Way to Build Good Habits & Break Bad Ones",
		Author:           "James Clear",
		Publisher:        "AVERY",
		Publication_Year: "2018",
		ISBN:             "9780735211292",
		NumberOfPage:     "356",
		Language:         "English",
	}
	if err := DB.Create(&bookInput).Error; err != nil {
		panic(err)
	}
	var createdBook models.Book
	DB.Last(&createdBook)
	return createdBook
}

func CleanSeeders() {
	DB.Exec("SET FOREIGN_KEY_CHECKS = 0")

	userResult := DB.Exec("DELETE FROM users")
	bookResult := DB.Exec("DELETE FROM books")

	var isFailed bool = bookResult.Error != nil || userResult.Error != nil 

	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}

	log.Println("Seeders are cleaned up successfully")
}

