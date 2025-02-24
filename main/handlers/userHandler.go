package handlers

import (
	dto "forum/main/dto/user"
	"forum/main/models"
	"forum/main/repositories"
	"forum/main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserRepository *repositories.UserRepository
}

func NewUserHandler(userRepository *repositories.UserRepository) *UserHandler {
	return &UserHandler{UserRepository: userRepository}
}

func (h *UserHandler) CreateUser(req *dto.CreateUserRequest) (*models.User, error) {
	hashedPassword, err := utils.HashString(req.Password)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash password")
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		IsActive: req.IsActive,
		Password: hashedPassword,
	}

	if err := h.UserRepository.CreateUser(&user); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	return &user, nil
}

func (h *UserHandler) GetAllUser() (*[]models.User, error) {
	users, err := h.UserRepository.GetAllUsers()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed list user")
	}
	return &users, nil
}

func (h *UserHandler) GetUserById(id string) (*models.User, error) {
	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	result, user := h.UserRepository.GetById(id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return user, nil
}

func (h *UserHandler) UpdateUser(id string, req *dto.UpdateUserRequest) (*models.User, error) {
	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	result, user := h.UserRepository.GetById(id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Age = req.Age

	if err := h.UserRepository.SaveUser(user); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	return user, nil
}

func (h *UserHandler) DeleteUser(id string) (map[string]string, error) {
	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	result, user := h.UserRepository.GetById(id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := h.UserRepository.SoftDeleteUser(user); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Fail to delete user")
	}

	return map[string]string{"message": "User deleted successfully"}, nil
}
