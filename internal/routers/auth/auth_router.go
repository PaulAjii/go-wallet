package auth

import (
	"github.com/PaulAjii/go-wallet/internal/handlers/auth"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(api fiber.Router, h *auth.AuthHandler) {
	auth := api.Group("/auth")

	// Staff management
	auth.Post("/register", h.Register)
	auth.Post("/login", h.Login)
	auth.Post("/verify-email", h.VerifyEmail)

}
