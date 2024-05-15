package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/charmbracelet/log"
	env "github.com/joho/godotenv"
)

func Init_PostGres() (*gorm.DB, error) {
	err := env.Load("../../.env")

	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err.Error())
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbSll := "disable"
	dbTimezone := "Asia/Ho_Chi_Minh"

	/* INFO: create dns string */
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		dbHost, dbUser, dbPass, dbName, dbPort, dbSll, dbTimezone)

	/* INFO: initialize charm logger */
	charmLogger := logger.New(log.New(os.Stdout), logger.Config{})

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: charmLogger,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err.Error())
	}

	return db, nil
}

// Define your database connect function below
