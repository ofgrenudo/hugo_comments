package common

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetConn() (*pgx.Conn, error) {
	// "postgres://${DB_USER}:${DB_PASSWORD}@db:${DB_PORT}/${DB_NAME}?sslmode=disable"
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	databaseURL := fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", db_user, db_password, db_port, db_name)

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	} else {
		defer conn.Close(context.Background())
	}
	return conn, nil
}
