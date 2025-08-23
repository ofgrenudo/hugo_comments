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
