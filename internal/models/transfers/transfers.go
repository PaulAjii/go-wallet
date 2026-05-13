package transfers

import "github.com/PaulAjii/go-wallet/internal/models"

type TransferModel struct {
	models.BaseWithoutUpdatedAt
	SenderWalletID   string `json:"senderWalletId" db:"sender_wallet_id"`
	ReceiverWalletID string `json:"receiverWalletId" db:"receiver_wallet_id"`
	Amount           int64  `json:"amount" db:"amount"`
	TransactionID    string `json:"transactionId" db:"transaction_id"`
}
