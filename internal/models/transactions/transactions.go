package transactions

import "github.com/PaulAjii/go-wallet/internal/models"

type TypeOfTransaction string

const (
	Credit TypeOfTransaction = "credit"
	Debit  TypeOfTransaction = "debit"
)

type TransactionModel struct {
	models.BaseWithoutUpdatedAt
	Reference     string                 `json:"reference" db:"reference"`
	WalletID      string                 `json:"walletId" db:"wallet_id"`
	Type          TypeOfTransaction      `json:"type" db:"type"`
	Category      models.Category        `json:"category" db:"category"`
	Amount        int64                  `json:"amount" db:"amount"`
	BalanceBefore int64                  `json:"balanceBefore" db:"balance_before"`
	BalanceAfter  int64                  `json:"balanceAfter" db:"balance_after"`
	Status        models.Status          `json:"status" db:"status"`
	Metadata      map[string]interface{} `json:"metadata,omitempty" db:"metadata"`
}
