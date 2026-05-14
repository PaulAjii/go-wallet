package auth

import (
	"github.com/PaulAjii/go-wallet/internal/handlers/auth"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(api fiber.Router, h *auth.AuthHandler) {
	staff := api.Group("/auth")

	// Staff management
	staff.Post("/register", h.Register)
	staff.Post("/login", h.Login)
}
