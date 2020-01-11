package models

type FlypayATransactions struct {
	Transactions []FlaypayATransaction `json:"transactions"`
}

type FlaypayATransaction struct {
	Amount         int    `json:"amount"`
	Currency       string `json:"currency"`
	StatusCode     int    `json:"statusCode"`
	OrderReference string `json:"orderReference"`
	TransactionId  string `json:"transactionId"`
}

var FlypayAStatusCodeMapping = map[int]string{
	1: "authorised",
	2: "decline",
	3: "refunded",
}

func (transaction FlaypayATransaction) GetAmount() int {
	return transaction.Amount
}

func (transaction FlaypayATransaction) GetCurrency() string {
	return transaction.Currency
}

func (transaction FlaypayATransaction) GetStatus() string {
	return FlypayAStatusCodeMapping[transaction.StatusCode]
}

func (transaction FlaypayATransaction) GetOrder() string {
	return transaction.OrderReference
}

func (transaction FlaypayATransaction) GetPayment() string {
	return transaction.TransactionId
}

func (transcation FlaypayATransaction) GetProvider() string {
	return "flypayA"
}
