package dtos

type FundWalletRequest struct {
	Amount      int64  `json:"amount"`
	CallbackURL string `json:"callbackUrl"`
}

type TransferRequest struct {
	Receiver  string `json:"receiver"`
	Amount    int64  `json:"amount"`
	Narration string `json:"narration"`
}

type WithdrawRequest struct {
	Amount      int64  `json:"amount"`
	BankCode    string `json:"bankCode"`
	AccountID   string `json:"accountID"`
	AccountName string `json:"accountName"`
}

type WalletResponse struct {
	AccountID string `json:"accountID"`
	Balance   int64  `json:"balance"`
	Currency  string `json:"currency"`
}
