package controllers

import (
	dto "forum/dto/post"
	"forum/handlers"
	"forum/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePostController(c echo.Context) error {
	req := c.Get("request").(*dto.CreatePostRequest)
	resp, err := handlers.CreatePost(req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func UpdatePostController(c echo.Context) error {
	id := c.QueryParam("id")
	req := c.Get("request").(*dto.UpdatePostRequest)
	resp, err := handlers.UpdatePost(id, req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func DeletePostController(c echo.Context) error {
	id := c.QueryParam("id")
	resp, err := handlers.DeletePost(id)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}
