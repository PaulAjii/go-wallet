package auth

import (
	"strings"

	"github.com/PaulAjii/go-wallet/internal/dtos"
	"github.com/PaulAjii/go-wallet/internal/usecases/auth"
	"github.com/PaulAjii/go-wallet/pkg/response"
	"github.com/PaulAjii/go-wallet/pkg/sysmsg"
	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	usecase *auth.AuthUsecase
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		usecase: auth.NewAuthCase(),
	}
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var payload dtos.RegisterRequest

	if err := c.Bind().Body(&payload); err != nil {
		return response.Error(c, sysmsg.ErrBadReq, err.Error(), fiber.StatusBadRequest)
	}

	user, err := h.usecase.Register(c.Context(), payload)
	if err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			return response.Error(c, sysmsg.ErrEmailAlreadyExists, err.Error(), fiber.StatusConflict)
		} else if strings.Contains(err.Error(), "username already exists") {
			return response.Error(c, sysmsg.ErrUsernameAlreadyExists, err.Error(), fiber.StatusConflict)
		} else if strings.Contains(err.Error(), "password length") {
			return response.Error(c, sysmsg.ErrInvalidPasswordLength, err.Error(), fiber.StatusBadRequest)
		} else {
			return response.Error(c, sysmsg.ErrInternalServerError, err.Error(), fiber.StatusInternalServerError)
		}
	}

	return response.JSON(c, sysmsg.CreationSuccess, user, fiber.StatusCreated)
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var payload dtos.LoginRequest

	if err := c.Bind().Body(&payload); err != nil {
		return response.Error(c, sysmsg.ErrBadReq, err.Error(), fiber.StatusBadRequest)
	}

	user, err := h.usecase.Login(c.Context(), payload)
	if err != nil {
		if strings.Contains(err.Error(), "invalid credentials") {
			return response.Error(c, sysmsg.ErrInvalidCredentials, err.Error(), fiber.StatusBadRequest)
		}

		return response.Error(c, sysmsg.ErrInternalServerError, err.Error(), fiber.StatusInternalServerError)
	}

	return response.JSON(c, sysmsg.LoginSuccess, user, fiber.StatusOK)
}
