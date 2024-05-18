package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	env "github.com/joho/godotenv"
)

func Init_PostGres() (gorm.Dialector, error) {
	err := env.Load("../../.env")

	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err.Error())
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbSll := "disable"
	dbTimezone := "Asia/Ho_Chi_Minh"

	/* INFO: create dns string */
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, dbSll, dbTimezone)

	dia := postgres.Open(dsn)

	return dia, nil
}

// Define your database connect function below
