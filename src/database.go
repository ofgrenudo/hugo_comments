package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

/** GetConn
*
* This function will attemtpt to create a connection to the database using the
* supplied credentials in the environment file. Without supplied variables, the
* program will crash with a fatal error.
*
* There are implicit expectations that you have the following variables declared
* in your .env file:
*
* - DB_USER
* - DB_PASS
* - DB_NAME
* - DB_PORT
* - DB_HOST
*
* @returns *pgx.Conn | nil
 */
func GetConn() (*pgx.Conn, error) {
	dbUser, dbUserExists := os.LookupEnv("DB_USER")
	dbPassword, dbPasswordExists := os.LookupEnv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	if !dbUserExists || !dbPasswordExists {
		log.Fatalf("Unable to find credentials in the environment file. Please update your environement variables and try again.")
	}

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return conn, nil
}

func EnsureCommentsTable(ctx context.Context, conn *pgx.Conn) error {

	createSchema := `
	CREATE SCHEMA IF NOT EXISTS hugo_comments
	`

	_, err := conn.Exec(ctx, createSchema)
	if err != nil {
		return fmt.Errorf("failed to create hugo_comments schema: %w", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS hugo_comments.comments (
		id SERIAL PRIMARY KEY,
		user_name TEXT NOT NULL,
		user_comment TEXT NOT NULL,
		user_ip TEXT,
		post_url TEXT NOT NULL,
		hidden BOOLEAN DEFAULT FALSE
	);
	`

	_, err = conn.Exec(ctx, createTable)
	if err != nil {
		return fmt.Errorf("failed to create comments table: %w", err)
	}

	return nil
}
