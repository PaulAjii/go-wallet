package paystack_references

import "github.com/PaulAjii/go-wallet/internal/models"

type PaystackReferenceModel struct {
	models.BaseWithoutUpdatedAt
	Reference            string          `json:"reference" db:"reference"`
	UserID               string          `json:"userId" db:"user_id"`
	Amount               int64           `json:"amount" db:"amount"`
	Type                 models.Category `json:"type" db:"type"`
	Status               models.Status   `json:"status" db:"status"`
	PaystackTransferCode string          `json:"paystackTransferCode" db:"paystack_transfer_code"`
}
