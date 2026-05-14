package users

import (
	"context"
	"errors"

	"github.com/PaulAjii/go-wallet/internal/models/users"
	"github.com/PaulAjii/go-wallet/pkg/database"
	"github.com/jackc/pgx/v5"
)

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (r *UsersRepository) CreateUsers(ctx context.Context, user *users.UsersModel) (*users.UsersModel, error) {
	stmt := `
		INSERT INTO users (full_name, username, email, password_hash, is_verified)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	var u users.UsersModel
	err := database.Pool.QueryRow(ctx, stmt,
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
