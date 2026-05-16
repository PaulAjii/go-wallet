package auth

import (
	"github.com/PaulAjii/go-wallet/internal/handlers/auth"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(api fiber.Router, h *auth.AuthHandler) {
	user := api.Group("/auth")

	// Staff management
	user.Post("/register", h.Register)
	user.Post("/login", h.Login)
	user.Post("/verify-email", h.VerifyEmail)

}
