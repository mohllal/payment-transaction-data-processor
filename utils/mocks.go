package utils

type TransactionsResponse struct {
	Status       string            `json:"status"`
	Transactions []TransactionFake `json:"result"`
}

type PingResponse struct {
	Message string `json:"message"`
}
