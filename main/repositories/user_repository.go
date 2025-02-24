package repositories

import (
	"forum/main/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := r.DB.Find(&users)
	return users, result.Error
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) SaveUser(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) GetById(id string) (*gorm.DB, *models.User) {
	var user models.User
	result := r.DB.Preload("Posts").First(&user, id)
	return result, &user
}

func (r *UserRepository) GetByEmail(email string) (*gorm.DB, *models.User) {
	var user models.User
	result := r.DB.Where("email = ?", email).First(&user)
	return result, &user
}

func (r *UserRepository) SoftDeleteUser(user *models.User) error {
	return r.DB.Delete(user).Error
}
