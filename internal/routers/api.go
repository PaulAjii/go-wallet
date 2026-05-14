package routers

import (
	authHandler "github.com/PaulAjii/go-wallet/internal/handlers/auth"
	"github.com/PaulAjii/go-wallet/internal/routers/auth"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// AUTH
	auth.SetupRoutes(api, authHandler.NewAuthHandler())
}
