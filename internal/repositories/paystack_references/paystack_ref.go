package paystack_references

import (
	"context"
	"errors"

	"github.com/PaulAjii/go-wallet/internal/models"
	"github.com/PaulAjii/go-wallet/internal/models/paystack_references"
	"github.com/PaulAjii/go-wallet/pkg/database"
	"github.com/jackc/pgx/v5"
)

type PaystackRefRepository struct{}

func NewPaystackRefRepository() *PaystackRefRepository {
	return &PaystackRefRepository{}
}

func (r *PaystackRefRepository) CreateReference(ctx context.Context, ref paystack_references.PaystackReferenceModel) error {
	stmt := `
		INSERT INTO paystack_references (reference, user_id, amount, type, status, paystack_transfer_code)
		VALUSES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	return database.Pool.QueryRow(
		ctx,
		stmt,
		ref.Reference,
		ref.UserID,
		ref.Amount,
		ref.Type,
		ref.Status,
		ref.PaystackTransferCode,
	).Scan(
		&ref.ID,
		&ref.CreatedAt,
	)
}

func (r *PaystackRefRepository) GetByReference(ctx context.Context, reference string) (*paystack_references.PaystackReferenceModel, error) {
	query := `
		SELECT id, reference, user_id, amount, type, status, paystack_transfer_code, created_at
		FROM paystack_references
		WHERE reference = $1
    `

	var ref paystack_references.PaystackReferenceModel
	err := database.Pool.QueryRow(ctx, query, reference).Scan(
		&ref.ID,
		&ref.Reference,
		&ref.UserID,
		&ref.Amount,
		&ref.Type,
		&ref.Status,
		&ref.PaystackTransferCode,
		&ref.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &ref, nil
}

func (r *PaystackRefRepository) UpdateStatus(ctx context.Context, reference string, status models.Status) error {
	query := `UPDATE paystack_references SET status = $1 WHERE reference = $2`
	_, err := database.Pool.Exec(ctx, query, status, reference)
	return err
}
