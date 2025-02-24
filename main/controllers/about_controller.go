package controllers

import (
	"forum/main/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AboutController(c echo.Context) error {
	res, err := handlers.AboutHandler()
	if err != nil {
		// Check if the error is an HTTPError
		if httpErr, ok := err.(*echo.HTTPError); ok {
			// Retrieve the status code and message from the HTTPError
			statusCode := httpErr.Code
			message := httpErr.Message

			// You can return a new error or handle it further
			return echo.NewHTTPError(statusCode, message)
		}
		// If it's not an HTTPError, handle it as a regular error
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, res)
}
