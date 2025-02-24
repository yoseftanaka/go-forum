package tests

import (
	"forum/main/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbout(t *testing.T) {
	expected := map[string]string{"message": "This is the About page."}

	// Call the function
	result, err := handlers.AboutHandler()

	// Assertions
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expected, result, "Response should match expected output")
}
