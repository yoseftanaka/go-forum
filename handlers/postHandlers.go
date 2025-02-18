package handlers

import (
	"forum/config"
	dto "forum/dto/post"
	"forum/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreatePost(c echo.Context) error {
	req := c.Get("request").(*dto.CreatePostRequest)

	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}

	if err := config.DB.Create(&post).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Fail to create post")
	}

	return c.JSON(http.StatusCreated, post)
}

func UpdatePost(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	var post models.Post
	result := config.DB.First(&post, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Post not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	req := c.Get("request").(*dto.UpdatePostRequest)

	post.Content = req.Content
	post.Title = req.Title
	post.IsEdited = true

	if err := config.DB.Save(&post).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Fail to update post")
	}

	return c.JSON(http.StatusOK, post)
}

func DeletePost(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	var post models.Post
	result := config.DB.First(&post, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Post not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Fail to delete post")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Post deleted successfully"})
}
