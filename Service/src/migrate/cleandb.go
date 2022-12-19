package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/basarrcan/NPCAI/initializers"
)

func main() {
	// Parse the command line flags
	dbName := flag.String("db", "", "Name of the database to purge")
	flag.Parse()

	// Check if a database name was provided
	if *dbName == "" {
		log.Fatal("You must provide a database name with the -db flag")
	}

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	// Connect to the PostgreSQL database
	conn := initializers.ConnectDB(&config)

	// Purge the database
	_, err = conn.Exec(context.Background(), fmt.Sprintf("DROP DATABASE %s", *dbName))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully purged database '%s'\n", *dbName)
}
