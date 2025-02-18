package routes

import "github.com/labstack/echo/v4"

// Register all routes
func RegisterRoutes(e *echo.Echo) {
	HomeRoutes(e)  // Calls home.go routes
	AboutRoutes(e) // Calls about.go routes
	UserRoutes(e)
	PostRoutes(e)
	AuthRoutes(e)
}
