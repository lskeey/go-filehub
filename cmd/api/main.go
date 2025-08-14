// File: cmd/api/main.go

package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lskeey/go-filehub/config"
	"github.com/lskeey/go-filehub/internal/database"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// 2. Connect to Database
	database.Connect(cfg)

	// 3. Initialize Gin Server
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong and connected to db",
		})
	})

	// 4. Run Server
	serverAddr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Server is running at %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
