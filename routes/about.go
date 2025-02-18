package routes

import (
	"forum/handlers"

	"github.com/labstack/echo/v4"
)

// Register About Route
func AboutRoutes(e *echo.Echo) {
	e.GET("/about", handlers.AboutHandler)
}
