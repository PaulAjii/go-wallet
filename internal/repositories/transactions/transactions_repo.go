package transactions

import (
	"context"
	"errors"

	tx "github.com/PaulAjii/go-wallet/internal/models/transactions"
	"github.com/PaulAjii/go-wallet/pkg/database"
	"github.com/jackc/pgx/v5"
)

type TransactionsRepository struct{}

func NewTransactionRepository() *TransactionsRepository {
	return &TransactionsRepository{}
}

func (r *TransactionsRepository) CreateTransaction(ctx context.Context, payload *tx.TransactionModel) (*tx.TransactionModel, error) {
	stmt := `
		INSERT INTO transactions (reference, wallet_id, type, category, amount, balance_before, balance_after, status, metadata)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, &9)
		RETURNING id, reference, wallet_id, type, category, amount, balance_before, balance_after, status, metadata, created_at
	`

	var t tx.TransactionModel
	err := database.Pool.QueryRow(
		ctx,
		stmt,
		payload.Reference,
		payload.WalletID,
		payload.Type,
		payload.Category,
		payload.Amount,
		payload.BalanceBefore,
		payload.BalanceAfter,
		payload.Status,
		payload.Metadata,
	).Scan(
		&t.ID,
		&t.Reference,
		&t.WalletID,
		&t.Type,
		&t.Category,
		&t.Amount,
		&t.BalanceBefore,
		&t.BalanceAfter,
		&t.Status,
		&t.Metadata,
		&t.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &t, nil
}

func (r *TransactionsRepository) GetByWalletID(ctx context.Context, walletID string, limit int) ([]*tx.TransactionModel, error) {
	query := `                                                                                                  
                SELECT id, reference, wallet_id, type, category, amount, balance_before, balance_after, status, created_at
                FROM transactions                                                                                   
                WHERE wallet_id = $1                                                                                  
                ORDER BY created_at DESC                                                                              
                LIMIT $2                                                                                              
        `

	rows, err := database.Pool.Query(ctx, query, walletID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var txns []*tx.TransactionModel
	for rows.Next() {
		var t tx.TransactionModel
		err := rows.Scan(
			&t.ID,
			&t.Reference,
			&t.WalletID,
			&t.Type,
			&t.Category,
			&t.Amount,
			&t.BalanceBefore,
			&t.BalanceAfter,
			&t.Status,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		txns = append(txns, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return txns, nil
}
