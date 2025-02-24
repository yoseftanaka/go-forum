package controllers

import (
	dto "forum/main/dto/post"
	"forum/main/handlers"
	"forum/main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	PostHandler *handlers.PostHandler
}

func NewPostController(postHandler *handlers.PostHandler) *PostController {
	return &PostController{PostHandler: postHandler}
}

func (pc *PostController) CreatePostController(c echo.Context) error {
	req := c.Get("request").(*dto.CreatePostRequest)
	resp, err := pc.PostHandler.CreatePost(req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func (pc *PostController) UpdatePostController(c echo.Context) error {
	id := c.QueryParam("id")
	req := c.Get("request").(*dto.UpdatePostRequest)
	resp, err := pc.PostHandler.UpdatePost(id, req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func (pc *PostController) DeletePostController(c echo.Context) error {
	id := c.QueryParam("id")
	resp, err := pc.PostHandler.DeletePost(id)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}
