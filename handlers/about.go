package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// About handler function
func AboutHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "This is the About page."})
}
