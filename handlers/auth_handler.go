package handlers

import (
	"fmt"
	"forum/config"
	"forum/constants"
	dto "forum/dto/user"
	"forum/models"
	"forum/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	err = config.RedisClient.Set(config.RedisContext, fmt.Sprintf("%s-%d", constants.USER, user.ID), token, 24*time.Hour).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
	}

	resp := &dto.LoginResponse{
		Token: token,
	}

	return resp, nil

}

func Logout(req *dto.LogoutRequest) (*map[string]string, error) {

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	err := config.RedisClient.Del(config.RedisContext, fmt.Sprintf("%s-%d", constants.USER, user.ID)).Err()
	if err != nil {
		fmt.Println("Error delete key:", err)
	}
	return &map[string]string{"message": "Logout successfully"}, nil
}
