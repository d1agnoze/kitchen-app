package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	env "github.com/joho/godotenv"
)

func Init_PostGres() (*gorm.DB, error) {
	err := env.Load()

	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err.Error())
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSll := os.Getenv("DB_GORM_SSL_MODE")
	dbTimezone := os.Getenv("DB_GORM_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, dbSll, dbTimezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Define your database connect function below
