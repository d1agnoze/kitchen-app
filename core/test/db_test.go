package test

import (
	"testing"

	"github.com/d1agnoze/kitchen/core/db"
	"github.com/joho/godotenv"
)

func TestDBConnection(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err)
	}
	/* INFO:   postgresql db connection test case */
	t.Run("Postgres DSN test", func(t *testing.T) {
		_, err := db.Init_PostGres()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Database connection test", func(t *testing.T) {
		_, err := db.ConnectDB()
		if err != nil {
			t.Error(err)
		}
	})
}
