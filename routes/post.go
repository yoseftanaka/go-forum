package routes

import (
	"forum/controllers"
	dto "forum/dto/post"
	"forum/middlewares"

	"github.com/labstack/echo/v4"
)

func PostRoutes(e *echo.Echo) {
	postGroup := e.Group("/posts")
	postGroup.Use(middlewares.JwtAuthMiddleware)
	postGroup.POST("/create", middlewares.RequestBinder(controllers.CreatePostController, &dto.CreatePostRequest{}))
	postGroup.PUT("/update", middlewares.RequestBinder(controllers.UpdatePostController, &dto.UpdatePostRequest{}))
	postGroup.DELETE("/delete", controllers.DeletePostController)
}
