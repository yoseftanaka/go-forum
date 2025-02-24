package routes

import (
	dto "forum/main/dto/user"
	"forum/main/loaders"
	"forum/main/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.Use(middlewares.JwtAuthMiddleware)
	userGroup.POST("/create", middlewares.RequestBinder(loaders.UserController.CreateUserCotnroller, &dto.CreateUserRequest{}))
	userGroup.GET("/get-list", loaders.UserController.GetAllUserCotnroller)
	userGroup.GET("/get-single", loaders.UserController.GetSingleUserCotnroller)
	userGroup.PUT("/update", middlewares.RequestBinder(loaders.UserController.UpdateUserCotnroller, &dto.UpdateUserRequest{}))
	userGroup.DELETE("/delete", loaders.UserController.DeleteUserCotnroller)
}
