package main

import (
	"context"
	"log"
)

func NewComment(username string, comment string, ip_address string, post_url string) {
	ctx := context.Background()
	conn, err := GetConn()
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer conn.Close(ctx)

	sqlStatement := `
		INSERT INTO hugo_comments.comments (user_name, user_comment, user_ip, post_url)	VALUES ($1, $2, $3, $4)`

	_, err = conn.Exec(context.Background(), sqlStatement, username, comment, ip_address, post_url)
	if err != nil {
		panic(err)
	}
}
