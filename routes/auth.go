package routes

import (
	dto "forum/dto/user"
	"forum/handlers"
	"forum/middlewares"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	e.POST("/login", middlewares.RequestBinder(handlers.Login, &dto.LoginRequest{}))
	e.POST("/logout", middlewares.RequestBinder(handlers.Logout, &dto.LogoutRequest{}))
}
