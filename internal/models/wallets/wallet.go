package wallets

import "github.com/PaulAjii/go-wallet/internal/models"

type WalletModel struct {
	models.BaseModel
	UserID    string `json:"userId" db:"user_id"`
	AccountID string `json:"accountID" db:"account_id"`
	Balance   int64  `json:"balance" db:"balance"`
	Currency  string `json:"currency" db:"currency"`
	IsActive  bool   `json:"isActive" db:"is_active"`
}
