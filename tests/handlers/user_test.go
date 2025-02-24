package tests

import (
	"errors"
	dto "forum/main/dto/user"
	"forum/main/handlers"
	"forum/main/models"
	"forum/main/utils"
	"forum/tests/mocks"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_Success(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userHandler := &handlers.UserHandler{UserRepository: mockUserRepo}
	utils.HashString = mocks.MockHashString

	req := &dto.CreateUserRequest{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Age:      25,
		IsActive: true,
		Password: "securepassword",
	}

	expectedUser := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		IsActive: req.IsActive,
		Password: "hashedPassword",
	}

	mockUserRepo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(nil)

	result, err := userHandler.CreateUser(req)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedUser.Name, result.Name)
	assert.Equal(t, expectedUser.Email, result.Email)
	assert.Equal(t, expectedUser.Password, result.Password)

	mockUserRepo.AssertExpectations(t)
}

func TestCreateUser_HashingError(t *testing.T) {
	// Setup
	mockUserRepo := new(mocks.MockUserRepository)
	userHandler := &handlers.UserHandler{UserRepository: mockUserRepo}
	utils.HashString = mocks.MockHashString

	req := &dto.CreateUserRequest{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Age:      25,
		IsActive: true,
		Password: "error", // Simulate hashing error
	}

	// Execute
	result, err := userHandler.CreateUser(req)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, http.StatusInternalServerError, err.(*echo.HTTPError).Code)
	assert.Equal(t, "Failed to hash password", err.(*echo.HTTPError).Message)
}

func TestCreateUser_RepoError(t *testing.T) {
	// Setup
	mockUserRepo := new(mocks.MockUserRepository)
	userHandler := &handlers.UserHandler{UserRepository: mockUserRepo}
	utils.HashString = mocks.MockHashString

	req := &dto.CreateUserRequest{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Age:      25,
		IsActive: true,
		Password: "securepassword",
	}

	mockUserRepo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(errors.New("database error"))

	// Execute
	result, err := userHandler.CreateUser(req)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, http.StatusInternalServerError, err.(*echo.HTTPError).Code)
	assert.Equal(t, "Failed to create user", err.(*echo.HTTPError).Message)

	mockUserRepo.AssertExpectations(t)
}
