package routes

import (
	dto "forum/dto/user"
	"forum/handlers"
	"forum/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.Use(middlewares.JwtAuthMiddleware)
	userGroup.POST("/create", middlewares.RequestBinder(handlers.CreateUser, &dto.CreateUserRequest{}))
	userGroup.GET("/get-list", handlers.GetAllUser)
	userGroup.GET("/get-single", handlers.GetUserById)
	userGroup.PUT("/update", middlewares.RequestBinder(handlers.UpdateUser, &dto.UpdateUserRequest{}))
	userGroup.DELETE("/delete", handlers.DeleteUser)
}
