package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrorResponse defines the standard error response structure
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// ErrorHandlerMiddleware is a global error handler
func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			// Convert to echo.HTTPError if possible
			if he, ok := err.(*echo.HTTPError); ok {
				return c.JSON(he.Code, ErrorResponse{
					Message: he.Message.(string),
					Status:  he.Code,
				})
			}

			// Default to internal server error if unknown
			return c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: "Internal Server Error",
				Status:  http.StatusInternalServerError,
			})
		}
		return nil
	}
}
