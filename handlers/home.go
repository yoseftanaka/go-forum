package handlers

import (
	"forum/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Home handler function
func HomeHandler(c echo.Context) error {
	resp := dto.HomeRs{
		Message: "Welcome to Home!",
		Status:  200,
	}
	return c.JSON(http.StatusOK, resp)
}
