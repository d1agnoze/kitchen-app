package test

import (
	"testing"

	"github.com/d1agnoze/kitchen/core/db"
	"github.com/d1agnoze/kitchen/services/auth/models"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func loadEnv() {
	godotenv.Load("../../.env")
}

func TestMigrations(t *testing.T) {
	loadEnv()
	db, err := db.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("User table migration test", func(t *testing.T) {
		// Check table for `Users` exists or not
		if !db.Migrator().HasTable("users") {
			t.Error("Table does not exist")
		}
	})
}

func TestSeedData(t *testing.T) {
	loadEnv()
	db, err := db.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	if !db.Migrator().HasTable("users") {
		t.Fatal("Table does not exist")
	}

	t.Run("User table seed test", func(t *testing.T) {
		var count int64
		db.Table("users").Count(&count)
		if count == 0 {
			t.Error("Table is empty")
		}
	})
	t.Run("Check if password hashing ok", func(t *testing.T) {
		// Check for table data exist or not
		var user models.User
		db.Table("users").Take(&user)

		if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte("123456")) != nil {
			t.Error("Password hashing check failed")
		}
	})
}
