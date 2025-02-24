package mocks

import (
	"forum/main/models"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

// CreateUser mocks the CreateUser method
func (m *MockUserRepository) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// GetAllUsers mocks fetching all users
func (m *MockUserRepository) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

// SaveUser mocks updating a user
func (m *MockUserRepository) SaveUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// GetById mocks fetching a user by ID
func (m *MockUserRepository) GetById(id string) (*gorm.DB, *models.User) {
	args := m.Called(id)
	return args.Get(0).(*gorm.DB), args.Get(1).(*models.User)
}

// GetByEmail mocks fetching a user by email
func (m *MockUserRepository) GetByEmail(email string) (*gorm.DB, *models.User) {
	args := m.Called(email)
	return args.Get(0).(*gorm.DB), args.Get(1).(*models.User)
}

// SoftDeleteUser mocks soft deleting a user
func (m *MockUserRepository) SoftDeleteUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}
