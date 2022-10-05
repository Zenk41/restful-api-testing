package routes

import (
	c "restful-api-testing/controllers"
	m "restful-api-testing/middlewares"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// Middleware
	m.LogMiddleware(e)
	// Login & Register
	e.POST("/register", c.Register)
	e.POST("/login", c.Login)
	// Authentication
	privateRoutes := e.Group("")
	privateRoutes.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte("secretToken"),
	}))

	// Checking Token
	privateRoutes.Use(m.CheckToken)

	// user routing
	// Not Authenticated
	e.POST("/users", c.CreateUserController) // Create User or Register
	// Authenticated
	privateRoutes.GET("/users", c.GetUsersController)
	privateRoutes.GET("/users/:id", c.GetUserController)
	privateRoutes.DELETE("/users/:id", c.DeleteUserController)
	privateRoutes.PUT("/users/:id", c.UpdateUserController)

	// book routing
	// Not Authenticated
	e.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)
	// Authenticated
	privateRoutes.POST("/books", c.CreateBookController)
	privateRoutes.DELETE("/books/:id", c.DeleteBookController)
	privateRoutes.PUT("/books/:id", c.UpdateBookController)

	return e
}
