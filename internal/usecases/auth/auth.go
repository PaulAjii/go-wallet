package auth

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"

	"github.com/PaulAjii/go-wallet/internal/dtos"
	"github.com/PaulAjii/go-wallet/internal/models/users"
	"github.com/PaulAjii/go-wallet/internal/models/wallets"
	usersRepo "github.com/PaulAjii/go-wallet/internal/repositories/users"
	walletsRepo "github.com/PaulAjii/go-wallet/internal/repositories/wallets"
	"github.com/PaulAjii/go-wallet/pkg/config"
	"github.com/PaulAjii/go-wallet/pkg/database"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	userRepo   *usersRepo.UsersRepository
	walletRepo *walletsRepo.WalletRepository
}

func NewAuthCase() *AuthUsecase {
	return &AuthUsecase{
		userRepo:   usersRepo.NewUsersRepository(),
		walletRepo: walletsRepo.NewWalletRepository(),
	}
}

func (a *AuthUsecase) Register(ctx context.Context, payload dtos.RegisterRequest) (*dtos.AuthResponse, error) {
	// Check uniqueness
	existing, err := a.userRepo.GetByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("email already exists")
	}

	existing, err = a.userRepo.GetByUsername(ctx, payload.Username)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("username already exists")
	}

	if utf8.RuneCountInString(payload.Password) < 8 {
		return nil, errors.New("password length must be equal to or more than 8")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	tx, err := database.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	user := users.UsersModel{
		FullName:   payload.FullName,
		Username:   payload.Username,
		Email:      payload.Email,
		Password:   string(hashed),
		IsVerified: false,
	}

	createdUser, err := a.userRepo.CreateUsers(ctx, tx, &user)
	if err != nil {
		return nil, err
	}

	wallet := wallets.WalletModel{
		UserID:    createdUser.ID.String(),
		AccountID: generateAccountID(),
		Balance:   0,
		Currency:  "NGN",
		IsActive:  true,
	}

	if err := a.walletRepo.CreateWallet(ctx, tx, &wallet); err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	token, err := generateToken(createdUser.ID.String())
	if err != nil {
		return nil, err
	}

	return &dtos.AuthResponse{
		Token: token,
		User: dtos.UserResponse{
			ID:         createdUser.ID.String(),
			FullName:   createdUser.FullName,
			Username:   createdUser.Username,
			Email:      createdUser.Email,
			IsVerified: createdUser.IsVerified,
		},
	}, nil
}

func (a *AuthUsecase) Login(ctx context.Context, payload dtos.LoginRequest) (*dtos.AuthResponse, error) {
	user, err := a.userRepo.GetByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := generateToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	return &dtos.AuthResponse{
		Token: token,
		User: dtos.UserResponse{
			ID:         user.ID.String(),
			FullName:   user.FullName,
			Username:   user.Username,
			Email:      user.Email,
			IsVerified: user.IsVerified,
		},
	}, nil
}

func generateAccountID() string {
	return fmt.Sprintf("%010d", 1000000000+rand.Intn(9000000000))
}

func generateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.ApplicationConfig.Supabase.JWTSecret))
}
