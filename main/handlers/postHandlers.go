package handlers

import (
	dto "forum/main/dto/post"
	"forum/main/models"
	"forum/main/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PostHandler struct {
	PostRepository repositories.IPostRepository
}

func NewPostHandler(postRepostiroy repositories.IPostRepository) *PostHandler {
	return &PostHandler{PostRepository: postRepostiroy}
}

func (h *PostHandler) CreatePost(req *dto.CreatePostRequest) (*models.Post, error) {
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}

	if err := h.PostRepository.CreatePost(&post); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Fail to create post")
	}

	return &post, nil
}

func (h *PostHandler) UpdatePost(id string, req *dto.UpdatePostRequest) (*models.Post, error) {
	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	result, post := h.PostRepository.GetById(id)

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

	if err := h.PostRepository.SavePost(post); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Fail to update post")
	}

	return post, nil
}

func (h *PostHandler) DeletePost(id string) (map[string]string, error) {
	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	result, post := h.PostRepository.GetById(id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "Post not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := h.PostRepository.SoftDeletePost(post); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Fail to delete post")
	}

	return map[string]string{"message": "Post deleted successfully"}, nil
}
