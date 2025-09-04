package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// https://gin-gonic.com/en/docs/examples/grouping-routes/``
	router := gin.Default()
	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		port = "8080"
	}

	// Generate Routes
	{
		v1 := router.Group("/api/v1")
		v1.GET("/health", ok_status)
	}

	// Initialize database
	check_comments_database()
	NewComment("ofgrenudo", "this is from a function", "loserville", "https://uhhhh")
	router.Run(":" + port)
}

func ok_status(c *gin.Context) {
	c.String(http.StatusOK, "200 OK")
}

func check_comments_database() {
	ctx := context.Background()

	conn, err := GetConn()
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer conn.Close(ctx)

	if err := EnsureCommentsTable(ctx, conn); err != nil {
		log.Fatalf("Error ensuring table: %v", err)
	}

	fmt.Println("comments table is ready")

}
