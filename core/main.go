package main

import (
	"flag"

	"github.com/charmbracelet/log"
	gormDB "github.com/d1agnoze/kitchen/core/db"
	env "github.com/joho/godotenv"
)

func init() {
	if err := env.Load("./.env"); err != nil {
		panic(err)
	}
}

func main() {
	seeder := flag.Bool("seed", false, "seed database")
	flag.Parse()

	db, err := gormDB.ConnectDB()

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err = gormDB.Migrate(db); err != nil {
		log.Fatalf("failed to migrate tables: %v", err)
	}

	switch {
	case *seeder:
		if err := gormDB.Seed(db); err != nil {
			log.Fatalf("failed to seed data: %v", err)
		}
	default:
		return
	}
}
