package handlers

import (
	"forum/config"
	dto "forum/dto/post"
	"forum/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreatePost(req *dto.CreatePostRequest) (*models.Post, error) {
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}

	if err := config.DB.Create(&post).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Fail to create post")
	}

	return &post, nil
}

func UpdatePost(id string, req *dto.UpdatePostRequest) (*models.Post, error) {
	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	var post models.Post
	result := config.DB.First(&post, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "Post not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	post.Content = req.Content
	post.Title = req.Title
	post.IsEdited = true

	if err := config.DB.Save(&post).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Fail to update post")
	}

	return &post, nil
}

func DeletePost(id string) (map[string]string, error) {
	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	var post models.Post
	result := config.DB.First(&post, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "Post not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Fail to delete post")
	}

	return map[string]string{"message": "Post deleted successfully"}, nil
}
