package routes

import (
	dto "forum/dto/post"
	"forum/handlers"
	"forum/middlewares"

	"github.com/labstack/echo/v4"
)

func PostRoutes(e *echo.Echo) {
	postGroup := e.Group("/posts")
	postGroup.Use(middlewares.JwtAuthMiddleware)
	postGroup.POST("/create", middlewares.RequestBinder(handlers.CreatePost, &dto.CreatePostRequest{}))
	postGroup.PUT("/update", middlewares.RequestBinder(handlers.UpdatePost, &dto.UpdatePostRequest{}))
	postGroup.DELETE("/delete", handlers.DeletePost)
}
