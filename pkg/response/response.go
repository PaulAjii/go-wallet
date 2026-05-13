package response

import "github.com/gofiber/fiber/v3"

type APIResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func JSON(c fiber.Ctx, message string, data interface{}, statusCode int) error

func Error(c fiber.Ctx, message string, err string, statusCode int) error
