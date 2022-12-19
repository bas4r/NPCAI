package services

import (
	"context"
	"fmt"
	"os"

	"github.com/basarrcan/NPCAI/models"
	"github.com/jackc/pgx/v4"
)

func ConnectDB(config *models.Config) *pgx.Conn {
	var err error
	ctx := context.Background()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)
	conn, err := pgx.Connect(ctx, connStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)
	return conn
}
