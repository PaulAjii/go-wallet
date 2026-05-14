package wallets

import (
	"context"
	"errors"

	"github.com/PaulAjii/go-wallet/internal/models/wallets"
	"github.com/PaulAjii/go-wallet/pkg/database"
	"github.com/jackc/pgx/v5"
)

type WalletRepository struct{}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{}
}

func (r *WalletRepository) CreateWallet(ctx context.Context, wallet wallets.WalletModel) error {
	stmt := `
		INSERT INTO wallets (user_id, account_id, balance, currency, is_active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	return database.Pool.QueryRow(ctx, stmt,
		wallet.UserID,
		wallet.AccountID,
		wallet.Balance,
		wallet.Currency,
		wallet.IsActive,
	).Scan(&wallet.ID, wallet.CreatedAt, wallet.UpdatedAt)
}

func (r *WalletRepository) GetByID(ctx context.Context, id string) (*wallets.WalletModel, error) {
	stmt := `
		SELECT id, user_id, account_id, balance, currency, is_active, created_at, updated_at
		FROM wallets
		WHERE id = $1
	`

	var w wallets.WalletModel
	err := database.Pool.QueryRow(ctx, stmt, id).Scan(
		&w.ID,
		&w.UserID,
		&w.AccountID,
		&w.Balance,
		&w.Currency,
		&w.IsActive,
		&w.CreatedAt,
		&w.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &w, nil
}

func (r *WalletRepository) GetByUserID(ctx context.Context, userID string) (*wallets.WalletModel, error) {
	stmt := `
		SELECT id, user_id, account_id, balance, currency, is_active, created_at, updated_at
		FROM wallets
		WHERE user_id = $1
	`

	var w wallets.WalletModel
	err := database.Pool.QueryRow(ctx, stmt, userID).Scan(
		&w.ID,
		&w.UserID,
		&w.AccountID,
		&w.Balance,
		&w.Currency,
		&w.IsActive,
		&w.CreatedAt,
		&w.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &w, nil
}

func (r *WalletRepository) GetByAccountID(ctx context.Context, accountID string) (*wallets.WalletModel, error) {
	stmt := `
		SELECT id, user_id, account_id, balance, currency, is_active, created_at, updated_at
		FROM wallets
		WHERE account_id = $1
	`

	var w wallets.WalletModel
	err := database.Pool.QueryRow(ctx, stmt, accountID).Scan(
		&w.ID,
		&w.UserID,
		&w.AccountID,
		&w.Balance,
		&w.Currency,
		&w.IsActive,
		&w.CreatedAt,
		&w.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &w, nil
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, walletID string, newBalance int64) error {
	query := `UPDATE wallets SET balance = $1 WHERE id = $2`
	_, err := database.Pool.Exec(ctx, query, newBalance, walletID)
	return err
}
