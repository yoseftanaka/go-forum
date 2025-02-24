package repositories

import (
	"forum/main/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostRepository) SavePost(post *models.Post) error {
	return r.DB.Save(post).Error
}

func (r *PostRepository) GetById(id string) (*gorm.DB, *models.Post) {
	var post models.Post
	result := r.DB.First(&post, id)
	return result, &post
}

func (r *PostRepository) SoftDeletePost(post *models.Post) error {
	return r.DB.Delete(post).Error
}
