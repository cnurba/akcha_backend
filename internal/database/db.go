package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type Database struct {
	Client *sql.DB
}

func NewDataBase() (*Database, error) {

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)

	dbConn, err := sql.Open("postgres", connectionString)

	if err != nil {
		return &Database{}, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Database{dbConn}, nil

}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.PingContext(ctx)
}
