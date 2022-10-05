package middlewares

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secretToken"))
}

func ExtractToken(e echo.Context) uint {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return uint(userId)
	}
	return 0
}

func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := ExtractToken(c)

		if userId == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid token",
			})
		}
		return next(c)
	}
}