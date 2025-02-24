package controllers

import (
	dto "forum/main/dto/user"
	"forum/main/handlers"
	"forum/main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthHandler *handlers.AuthHandler
}

func NewAuthController(authHandler *handlers.AuthHandler) *AuthController {
	return &AuthController{AuthHandler: authHandler}
}

func (ac *AuthController) LoginController(c echo.Context) error {
	req := c.Get("request").(*dto.LoginRequest)
	resp, err := ac.AuthHandler.Login(req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func (ac *AuthController) LogoutController(c echo.Context) error {
	req := c.Get("request").(*dto.LogoutRequest)
	resp, err := ac.AuthHandler.Logout(req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}
