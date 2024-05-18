package db

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/charmbracelet/log"
)

func Init_Database() (*gorm.DB, error) {
  // change your preferred database here
	dialector, err := Init_PostGres()

	charmLogger := logger.New(log.New(os.Stdout), logger.Config{})

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: charmLogger,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err.Error())
	}

	return db, nil
}
