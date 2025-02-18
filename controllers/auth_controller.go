package controllers

import (
	dto "forum/dto/user"
	"forum/handlers"
	"forum/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	req := c.Get("request").(*dto.LoginRequest)
	resp, err := handlers.Login(req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func LogoutController(c echo.Context) error {
	req := c.Get("request").(*dto.LogoutRequest)
	resp, err := handlers.Logout(req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}
