package routes

import (
	dto "forum/main/dto/user"
	"forum/main/loaders"
	"forum/main/middlewares"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	e.POST("/login", middlewares.RequestBinder(loaders.AuthController.LoginController, &dto.LoginRequest{}))
	e.POST("/logout", middlewares.RequestBinder(loaders.AuthController.LogoutController, &dto.LogoutRequest{}))
}
