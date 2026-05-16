package users

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/PaulAjii/go-wallet/internal/dtos"
	"github.com/PaulAjii/go-wallet/internal/models/users"
	"github.com/PaulAjii/go-wallet/pkg/database"
	"github.com/jackc/pgx/v5"
)

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (r *UsersRepository) CreateUsers(ctx context.Context, q database.Querier, user *users.UsersModel) (*users.UsersModel, error) {
	stmt := `
		INSERT INTO users (full_name, username, email, password_hash, is_verified)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, full_name, username, email, is_verified, created_at, updated_at
	`

	var u users.UsersModel
	err := q.QueryRow(ctx, stmt,
		&user.FullName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.IsVerified,
	).Scan(
		&u.ID,
		&u.FullName,
		&u.Username,
		&u.Email,
		&u.IsVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *UsersRepository) GetByID(ctx context.Context, id string) (*users.UsersModel, error) {
	stmt := `
		SELECT id, full_name, username, email, is_verified, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var u users.UsersModel
	err := database.Pool.QueryRow(ctx, stmt, id).Scan(
		&u.ID,
		&u.FullName,
		&u.Username,
		&u.Email,
		&u.IsVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *UsersRepository) GetByEmail(ctx context.Context, email string) (*users.UsersModel, error) {
	stmt := `
		SELECT id, full_name, username, password_hash, email, is_verified, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var u users.UsersModel
	err := database.Pool.QueryRow(ctx, stmt, email).Scan(
		&u.ID,
		&u.FullName,
		&u.Username,
		&u.Password,
		&u.Email,
		&u.IsVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *UsersRepository) GetByUsername(ctx context.Context, username string) (*users.UsersModel, error) {
	stmt := `
		SELECT id, full_name, username, password_hash, email, is_verified, created_at, updated_at
		FROM users
		WHERE username = $1
	`

	var u users.UsersModel
	err := database.Pool.QueryRow(ctx, stmt, username).Scan(
		&u.ID,
		&u.FullName,
		&u.Username,
		&u.Password,
		&u.Email,
		&u.IsVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *UsersRepository) Verify(ctx context.Context, id string) error {
	stmt := `
		UPDATE users SET is_verified = true WHERE id = $1
	`

	_, err := database.Pool.Exec(ctx, stmt, id)
	return err
}

func (r *UsersRepository) GetProfile(ctx context.Context, id string) (*dtos.ProfileDTO, error) {
	stmt := `
        SELECT 
            u.id, u.full_name, u.username, u.email, u.is_verified,
            w.account_number, w.balance, w.currency,
            COALESCE(
                (
                    SELECT jsonb_agg(
                        jsonb_build_object(
                            'reference', t.reference,
                            'type', t.type,
                            'category', t.category,
                            'amount', t.amount,
                            'balanceBefore', t.balance_before,
                            'balanceAfter', t.balance_after,
                            'status', t.status,
                            'createdAt', t.created_at
                        ) ORDER BY t.created_at DESC
                    )
                    FROM (
                        SELECT * FROM transactions
                        WHERE wallet_id = w.id
                        ORDER BY created_at DESC
                        LIMIT 10
                    ) t
                ),
                '[]'::jsonb
            ) AS transactions
        FROM users u
        JOIN wallets w ON w.user_id = u.id
        WHERE u.id = $1
    `

	var (
		txsJSON  []byte
		userJSON dtos.ProfileDTO
	)

	err := database.Pool.QueryRow(ctx, stmt, id).Scan(
		&userJSON.User.ID,
		&userJSON.User.FullName,
		&userJSON.User.Username,
		&userJSON.User.Email,
		&userJSON.User.IsVerified,
		&userJSON.Wallet.AccountNumber,
		&userJSON.Wallet.Balance,
		&userJSON.Wallet.Currency,
		&txsJSON,
	)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(txsJSON, &userJSON.Transaction); err != nil {
		return nil, err
	}

	return &userJSON, nil
}
