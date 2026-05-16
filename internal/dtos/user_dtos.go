package dtos

type ProfileDTO struct {
	User        UserResponse          `json:"user_info"`
	Wallet      WalletResponse        `json:"wallet"`
	Transaction []TransactionResponse `json:"transactions"`
}
