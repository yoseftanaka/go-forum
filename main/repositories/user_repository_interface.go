package repositories

import (
	"forum/main/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) error
	SaveUser(user *models.User) error
	GetById(id string) (*gorm.DB, *models.User)
	GetByEmail(email string) (*gorm.DB, *models.User)
	SoftDeleteUser(user *models.User) error
}
