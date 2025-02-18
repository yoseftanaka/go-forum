package tests

import (
	"testing"
)

func TestAbout(t *testing.T) {
	// e := echo.New()

	// // Mock the response recorder
	// rec := httptest.NewRecorder()

	// // Create a mock Echo context
	// req := httptest.NewRequest(http.MethodGet, "/about", nil) // Create the request
	// c := e.NewContext(req, rec)                               // Initialize the context with the request and response recorder

	// // Call the AboutHandler directly
	// err := handlers.AboutHandler(c)

	// // Assert no errors
	// assert.NoError(t, err)

	// // Check the response status code
	// assert.Equal(t, http.StatusOK, rec.Code)

	// // Check the response body
	// expectedResponse := `{"message":"This is the About page."}`
	// actualResponse := strings.TrimSpace(rec.Body.String()) // Trim any extra spaces/newlines
	// assert.Equal(t, expectedResponse, actualResponse)
}
