package users

import (
	"github.com/PaulAjii/go-wallet/internal/handlers/users"
	"github.com/PaulAjii/go-wallet/pkg/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(api fiber.Router, h *users.UserHandler) {
	users := api.Group("/users")
	users.Use(middleware.AuthMiddleware())

	// Staff management
	users.Get("/me", h.GetProfile)

}
