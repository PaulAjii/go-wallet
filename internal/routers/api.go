package routers

import (
	authHandler "github.com/PaulAjii/go-wallet/internal/handlers/auth"
	usersHandler "github.com/PaulAjii/go-wallet/internal/handlers/users"
	"github.com/PaulAjii/go-wallet/internal/routers/auth"
	"github.com/PaulAjii/go-wallet/internal/routers/users"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// AUTH
	auth.SetupRoutes(api, authHandler.NewAuthHandler())

	// USERS
	users.SetupRoutes(api, usersHandler.NewUserHandler())
}
