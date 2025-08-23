package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	conn, err := GetConn()
	if err != nil {
		log.Fatalf("Unable to connect to the database. Please confirm the database connection paramaters %s.", err)
	}

	defer conn.Close(context.Background())

	fmt.Println("Successfully connected to the database.")
}
