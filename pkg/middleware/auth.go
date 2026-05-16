package middleware

import (
	"strings"

	"github.com/PaulAjii/go-wallet/pkg/config"
	"github.com/PaulAjii/go-wallet/pkg/response"
	"github.com/PaulAjii/go-wallet/pkg/sysmsg"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Error(c, sysmsg.ErrUnauthorized, "missing authorization header", fiber.StatusUnauthorized)
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return response.Error(c, sysmsg.ErrUnauthorized, "invalid authorization token", fiber.StatusUnauthorized)
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.ApplicationConfig.Supabase.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			return response.Error(c, sysmsg.ErrUnauthorized, "invalid or expired token", fiber.StatusUnauthorized)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return response.Error(c, sysmsg.ErrUnauthorized, "invalid token claims", fiber.StatusUnauthorized)
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			return response.Error(c, sysmsg.ErrUnauthorized, "invalid user id", fiber.StatusUnauthorized)
		}

		c.Locals("userID", sub)
		return c.Next()
	}
}

func GetUserID(c fiber.Ctx) string {
	return c.Locals("userID").(string)
}
