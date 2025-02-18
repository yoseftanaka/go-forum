package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load the main .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No default .env file found")
	}

	// Get the ENV variable
	env := os.Getenv("ENV")
	if env == "" {
		env = "local" // Default to local if ENV is not set
	}

	// Load environment-specific .env file
	envFile := fmt.Sprintf(".env-%s", env)
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("No %s file found, using default .env\n", envFile)
	}
}
