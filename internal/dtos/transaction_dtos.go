package dtos

type TransactionResponse struct {
	Reference     string `json:"reference"`
	Type          string `json:"type"`
	Category      string `json:"category"`
	Amount        int64  `json:"amount"`
	BalanceBefore int64  `json:"balanceBefore"`
	BalanceAfter  int64  `json:"balanceAfter"`
	Status        string `json:"status"`
	CreatedAt     string `json:"createdAt"`
}
