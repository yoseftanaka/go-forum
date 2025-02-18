package controllers

import (
	dto "forum/dto/user"
	"forum/handlers"
	"forum/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserCotnroller(c echo.Context) error {
	req := c.Get("request").(*dto.CreateUserRequest)
	resp, err := handlers.CreateUser(req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func GetAllUserCotnroller(c echo.Context) error {
	resp, err := handlers.GetAllUser()

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func GetSingleUserCotnroller(c echo.Context) error {
	id := c.QueryParam("id")
	resp, err := handlers.GetUserById(id)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func UpdateUserCotnroller(c echo.Context) error {
	id := c.QueryParam("id")
	req := c.Get("request").(*dto.UpdateUserRequest)

	resp, err := handlers.UpdateUser(id, req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func DeleteUserCotnroller(c echo.Context) error {
	id := c.QueryParam("id")
	resp, err := handlers.DeleteUser(id)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}
