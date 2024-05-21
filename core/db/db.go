package db

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/charmbracelet/log"
	"github.com/d1agnoze/kitchen/core/seed"
	m "github.com/d1agnoze/kitchen/services/auth/models"
)

func ConnectDB() (*gorm.DB, error) {
	// change your preferred database here
	dialector, err := Init_PostGres()
	if err != nil {
		return nil, fmt.Errorf("failed to init dns: %v", err.Error())
	}
	charmLogger := logger.New(log.New(os.Stdout), logger.Config{})

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: charmLogger,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err.Error())
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&m.User{})
	if err != nil {
		return err
	}
	return nil
}

func Seed(db *gorm.DB) error {
	if err := seed.CreateUserSeeds(db, seed.Options{Count: 10}); err != nil {
		return err
	}
	return nil
}
