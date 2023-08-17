package entity

type Request struct {
	СurrencyСode string  `json:"сurrency_сode"`
	Amount       float64 `json:"amount"`
	NumberWallet int     `json:"number_wallet"`
}

type Response struct {
	СurrencyСode string  `json:"сurrency_сode"`
	Amount       float64 `json:"amount"`
	NumberWallet int     `json:"number_wallet"`
	Status       string  `json:"status"`
}

type CheckWalletRequest struct {
	NumberWallet int `json:"number_wallet"`
}

type CheckWalletResponse struct {
	NumberWallet int    `json:"number_wallet"`
	Status       string `json:"status"`
}
