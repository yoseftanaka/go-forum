package routes

import (
	"forum/controllers"

	"github.com/labstack/echo/v4"
)

// Register About Route
func AboutRoutes(e *echo.Echo) {
	e.GET("/about", controllers.AboutController)
}
