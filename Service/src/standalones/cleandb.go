package main

import (
	"fmt"
	"os"

	"github.com/basarrcan/NPCAI/services"
)

func main() {
	// Connect to the PostgreSQL database
	db := services.ConnectDB()

	dbName := os.Args[1]

	// Purge the database using gorm
	// db.Exec(fmt.Sprintf("DROP DATABASE %s", dbName))
	db.Migrator().DropTable(dbName)

	fmt.Printf("Successfully purged database '%s'\n", dbName)
}
