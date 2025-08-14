// File: cmd/api/main.go

package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lskeey/go-filehub/config"
	"github.com/lskeey/go-filehub/internal/database"
	"github.com/lskeey/go-filehub/internal/handler"
	"github.com/lskeey/go-filehub/internal/repository"
	"github.com/lskeey/go-filehub/internal/service"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// 2. Connect to Database
	database.Connect(cfg)
	db := database.DB

	// 3. Initialize Layers (Dependency Injection)
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, cfg)
	authHandler := handler.NewAuthHandler(authService)

	// 4. Initialize Gin Server
	r := gin.Default()

	// 5. Setup Routes
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 6. Run Server
	serverAddr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Server is running at %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
