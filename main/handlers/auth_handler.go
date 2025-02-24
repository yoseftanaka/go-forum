package handlers

import (
	"fmt"
	"forum/main/config"
	"forum/main/constants"
	dto "forum/main/dto/user"
	"forum/main/repositories"
	"forum/main/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	UserRepository *repositories.UserRepository
}

func NewAuthHandler(userRepository *repositories.UserRepository) *AuthHandler {
	return &AuthHandler{UserRepository: userRepository}
}

func (h *AuthHandler) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {

	result, user := h.UserRepository.GetByEmail(req.Email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "Invalid email or password")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
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

func (h *AuthHandler) Logout(req *dto.LogoutRequest) (*map[string]string, error) {

	result, user := h.UserRepository.GetByEmail(req.Email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "Invalid email or password")
		}
		// Handle other types of errors (e.g., database connectivity errors)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	err := config.RedisClient.Del(config.RedisContext, fmt.Sprintf("%s-%d", constants.USER, user.ID)).Err()
	if err != nil {
		fmt.Println("Error delete key:", err)
	}
	return &map[string]string{"message": "Logout successfully"}, nil
}
