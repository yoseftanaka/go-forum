package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RequestBinder is a middleware to bind and validate request data
func RequestBinder(next echo.HandlerFunc, model interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Create a new instance of the model
		req := model

		// Bind request body to struct
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Validate request
		// if err := c.Validate(req); err != nil {
		// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		// }

		// Store parsed data in context
		c.Set("request", req)

		// Call the next handler
		return next(c)
	}
}
