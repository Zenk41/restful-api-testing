package repository

import (
	"restful-api-testing/database"
	mid "restful-api-testing/middlewares"
	"restful-api-testing/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepo struct{}

func (a *AuthRepo) Register(input models.UserInput) (*models.User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser models.User = models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(password),
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (a *AuthRepo) Login(input models.UserLogin) string {
	var user models.User = models.User{}

	if err := database.DB.First(&user, "email=?", input.Email).Error; err != nil {
		return ""
	}
	if user.ID == 0 {
		return ""
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return ""
	}
	token, err := mid.CreateToken(user.ID)
	if err != nil {
		return ""
	}
	return token
}
