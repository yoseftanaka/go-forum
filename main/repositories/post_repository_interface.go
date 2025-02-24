package repositories

import (
	"forum/main/models"

	"gorm.io/gorm"
)

type IPostRepository interface {
	CreatePost(post *models.Post) error
	SavePost(post *models.Post) error
	GetById(id string) (*gorm.DB, *models.Post)
	SoftDeletePost(post *models.Post) error
}
