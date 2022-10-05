package services

import (
	"restful-api-testing/database"
	"restful-api-testing/models"
)

// get all users
func GetUsers() (interface{}, error) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// get user by id
func GetUserById(id int) (interface{}, error) {
	var user []models.User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// create new user
// func CreateUser(user models.User) (interface{}, error) {
// 	if err := database.DB.Create(&user).Error; err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// } Changing Createuser to Register

// update user by id
func UpdateUser(id int, user models.User) (interface{}, error) {
	users := models.User{}
	if err := database.DB.Model(&users).Where("id=?", id).Updates(&user).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// delete user by id
func DeleteUser(id int) (interface{}, error) {
	var user []models.User
	if err := database.DB.Where("id=?", id).First(&user).Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
