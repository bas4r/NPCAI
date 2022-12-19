package migrate

import (
	"fmt"
	"log"

	"github.com/basarrcan/NPCAI/initializers"
	"github.com/basarrcan/NPCAI/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}


package main

import (
	"context"
	"flag"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v4"
)

func main() {
	// Parse the command line flags
	migrationPath := flag.String("path", "", "Path to the migration files")
	flag.Parse()

	// Check if a migration path was provided
	if *migrationPath == "" {
		log.Fatal("You must provide a migration path with the -path flag")
	}

	// Connect to the PostgreSQL database
	conn, err := pgx.Connect(context.Background(), "postgresql://user:password@localhost/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Set up the migrate library
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", *migrationPath),
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	// Perform the migration
	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migration complete")
}
