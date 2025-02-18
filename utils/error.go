package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleError handles error responses and returns an appropriate HTTPError
func HandleError(err error) *echo.HTTPError {
	if err != nil {
		// Check if the error is an HTTPError
		if httpErr, ok := err.(*echo.HTTPError); ok {
			// Retrieve the status code and message from the HTTPError
			statusCode := httpErr.Code
			message := httpErr.Message
			// Return the same HTTPError
			return echo.NewHTTPError(statusCode, message)
		}
		// If it's not an HTTPError, handle it as a regular error
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return nil
}
