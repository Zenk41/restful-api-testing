package controllers

import (
	"net/http"
	"restful-api-testing/models"
	"restful-api-testing/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserResponse struct {
	Message string
	User    []models.UserInput
}

// get all users
func GetUsersController(c echo.Context) error {
	users, err := services.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	var UserInput *models.UserInput = new(models.UserInput)
	if err := c.Bind(UserInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request",
		})
	}
	if err := UserInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "validation failed",
		})
	}
	users, err := authService.Register(*UserInput)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success creating new user",
		"user":    users,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	users := models.User{}
	if err := c.Bind(&users); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.UpdateUser(id, users)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}
