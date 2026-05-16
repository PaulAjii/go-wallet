package users

import (
	"github.com/PaulAjii/go-wallet/internal/usecases/users"
	"github.com/PaulAjii/go-wallet/pkg/middleware"
	"github.com/PaulAjii/go-wallet/pkg/response"
	"github.com/PaulAjii/go-wallet/pkg/sysmsg"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	u *users.UserUsecase
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		u: users.NewUserUsecase(),
	}
}

func (h *UserHandler) GetProfile(c fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	profile, err := h.u.GetProfile(c.Context(), userID)
	if err != nil {
		return response.Error(c, sysmsg.ErrFailedToLoadProfile, err.Error(), fiber.StatusInternalServerError)
	}

	return response.JSON(c, sysmsg.UserProfileFetchSuccess, profile, fiber.StatusOK)
}
