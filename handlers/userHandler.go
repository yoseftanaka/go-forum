package handlers

import (
	"forum/config"
	dto "forum/dto/user"
	"forum/models"
	"forum/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateUser(req *dto.CreateUserRequest) (*models.User, error) {
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

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	return &user, nil
}

func GetAllUser() (*[]models.User, error) {
	var users []models.User
	config.DB.Where(&models.User{IsActive: true}).Find(&users)
	return &users, nil
}

func GetUserById(id string) (*models.User, error) {
	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	var user models.User
	result := config.DB.Preload("Posts").First(&user, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return &user, nil
}

func UpdateUser(id string, req *dto.UpdateUserRequest) (*models.User, error) {
	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	var user models.User
	result := config.DB.First(&user, id)

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

	if err := config.DB.Save(&user).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Fail to update user")
	}

	return &user, nil
}

func DeleteUser(id string) (map[string]string, error) {
	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	var user models.User
	result := config.DB.First(&user, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Fail to delete user")
	}

	return map[string]string{"message": "User deleted successfully"}, nil
}
