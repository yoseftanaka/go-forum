package routes

import (
	dto "forum/main/dto/post"
	"forum/main/loaders"
	"forum/main/middlewares"

	"github.com/labstack/echo/v4"
)

func PostRoutes(e *echo.Echo) {
	postGroup := e.Group("/posts")
	postGroup.Use(middlewares.JwtAuthMiddleware)
	postGroup.POST("/create", middlewares.RequestBinder(loaders.PostController.CreatePostController, &dto.CreatePostRequest{}))
	postGroup.PUT("/update", middlewares.RequestBinder(loaders.PostController.UpdatePostController, &dto.UpdatePostRequest{}))
	postGroup.DELETE("/delete", loaders.PostController.DeletePostController)
}
