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

func CreateUser(c echo.Context) error {
	req := c.Get("request").(*dto.CreateUserRequest)

	hashedPassword, err := utils.HashString(req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash password")
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		IsActive: req.IsActive,
		Password: hashedPassword,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, user)
}

func GetAllUser(c echo.Context) error {
	var users []models.User
	config.DB.Where(&models.User{IsActive: true}).Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUserById(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	var user models.User
	result := config.DB.Preload("Posts").First(&user, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	var user models.User
	result := config.DB.First(&user, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	req := c.Get("request").(*dto.UpdateUserRequest)

	user.Name = req.Name
	user.Email = req.Email
	user.Age = req.Age

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Fail to update user")
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "ID not defined")
	}

	var user models.User
	result := config.DB.First(&user, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Fail to delete user")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
