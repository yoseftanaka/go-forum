package routes

import (
	"forum/main/handlers"

	"github.com/labstack/echo/v4"
)

// Register Home Route
func HomeRoutes(e *echo.Echo) {
	e.GET("/", handlers.HomeHandler)
}
