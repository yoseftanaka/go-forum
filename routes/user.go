package routes

import (
	"forum/controllers"
	dto "forum/dto/user"
	"forum/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.Use(middlewares.JwtAuthMiddleware)
	userGroup.POST("/create", middlewares.RequestBinder(controllers.CreateUserCotnroller, &dto.CreateUserRequest{}))
	userGroup.GET("/get-list", controllers.GetAllUserCotnroller)
	userGroup.GET("/get-single", controllers.GetSingleUserCotnroller)
	userGroup.PUT("/update", middlewares.RequestBinder(controllers.UpdateUserCotnroller, &dto.UpdateUserRequest{}))
	userGroup.DELETE("/delete", controllers.DeleteUserCotnroller)
}
