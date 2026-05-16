package users

import (
	"context"

	"github.com/PaulAjii/go-wallet/internal/dtos"
	"github.com/PaulAjii/go-wallet/internal/repositories/users"
)

type UserUsecase struct {
	uRepo *users.UsersRepository
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{
		uRepo: users.NewUsersRepository(),
	}
}

func (u *UserUsecase) GetProfile(ctx context.Context, id string) (*dtos.ProfileDTO, error) {
	return u.uRepo.GetProfile(ctx, id)
}
