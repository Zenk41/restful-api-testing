package controllers

import (
	"net/http"
	"restful-api-testing/models"
	"restful-api-testing/services"

	"github.com/labstack/echo/v4"
)

var authService services.AuthService = services.NewAuthService()

func Register(c echo.Context) error {
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

	user, err := authService.Register(*UserInput)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success register new user",
		"user":    user,
	})
}

func Login(c echo.Context) error {
	var UserInput *models.UserLogin = new(models.UserLogin)

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
	token := authService.Login(*UserInput)
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "invalid email or password",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
