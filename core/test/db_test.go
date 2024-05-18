package test

import (
	"testing"

	"github.com/d1agnoze/kitchen/core/db"
)

func TestDBConnection(t *testing.T) {
	/* INFO:   postgresql db connection test case */
	t.Run("Postgres DSN test", func(t *testing.T) {
		_, err := db.Init_PostGres()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Database connection test", func(t *testing.T) {
		_, err := db.Init_Database()
		if err != nil {
			t.Error(err)
		}
	})
}
