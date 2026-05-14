package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/PaulAjii/go-wallet/internal/routers"
	"github.com/PaulAjii/go-wallet/pkg/config"
	"github.com/PaulAjii/go-wallet/pkg/database"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	config.Load()

	database.Connect()
	defer database.Close()

	app := fiber.New(
		fiber.Config{
			AppName: "Go Wallet",
		},
	)

	allowedOrigins := config.ApplicationConfig.App.AllowedOrigins
	if allowedOrigins == "" {
		allowedOrigins = "*" // Fallback to * for local dev if not specified
	}

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{allowedOrigins},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: false,
	}))

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "Go Wallet is healthy",
		})
	})

	// Setup Routes
	routers.SetupRoutes(app)

	serverErrors := make(chan error, 1)

	go func() {
		port := config.ApplicationConfig.App.Port
		if port == "" {
			port = "8080"
		}
		log.Printf("Go Wallet is running on port %s", port)
		serverErrors <- app.Listen(":" + port)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Printf("Error starting server: %v", err)
	case sig := <-quit:
		log.Printf("Received signal: %v. Shutting down server...", sig)
		if err := app.Shutdown(); err != nil {
			log.Printf("Error during server shutdown: %v", err)
		}
	}

	log.Println("Server gracefully stopped")
}
