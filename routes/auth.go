package routes

import (
	"forum/controllers"
	dto "forum/dto/user"
	"forum/middlewares"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	e.POST("/login", middlewares.RequestBinder(controllers.LoginController, &dto.LoginRequest{}))
	e.POST("/logout", middlewares.RequestBinder(controllers.LogoutController, &dto.LogoutRequest{}))
}
