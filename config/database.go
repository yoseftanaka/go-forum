package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare DB variable globally
var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	host := os.Getenv("db_host")
	user := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	port := os.Getenv("db_port")
	sslMode := os.Getenv("db_sslmode")
	timezone := os.Getenv("db_timezone")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=" + sslMode + " TimeZone=" + timezone
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")

	err = database.AutoMigrate(Models...)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
}
