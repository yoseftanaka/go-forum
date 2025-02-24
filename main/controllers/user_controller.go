package controllers

import (
	dto "forum/main/dto/user"
	"forum/main/handlers"
	"forum/main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserHandler *handlers.UserHandler
}

func NewUserController(userHandler *handlers.UserHandler) *UserController {
	return &UserController{UserHandler: userHandler}
}

func (uc *UserController) CreateUserCotnroller(c echo.Context) error {
	req := c.Get("request").(*dto.CreateUserRequest)
	resp, err := uc.UserHandler.CreateUser(req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func (uc *UserController) GetAllUserCotnroller(c echo.Context) error {
	resp, err := uc.UserHandler.GetAllUser()

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func (uc *UserController) GetSingleUserCotnroller(c echo.Context) error {
	id := c.QueryParam("id")
	resp, err := uc.UserHandler.GetUserById(id)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func (uc *UserController) UpdateUserCotnroller(c echo.Context) error {
	id := c.QueryParam("id")
	req := c.Get("request").(*dto.UpdateUserRequest)

	resp, err := uc.UserHandler.UpdateUser(id, req)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}

func (uc *UserController) DeleteUserCotnroller(c echo.Context) error {
	id := c.QueryParam("id")
	resp, err := uc.UserHandler.DeleteUser(id)

	if httpErr := utils.HandleError(err); httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, resp)
}
