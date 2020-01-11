package models

type FlypayBTransactions struct {
	Transactions []FlaypayBTransaction `json:"transactions"`
}

type FlaypayBTransaction struct {
	Value               int    `json:"value"`
	TransactionCurrency string `json:"transactionCurrency"`
	StatusCode          int    `json:"statusCode"`
	OrderInfo           string `json:"orderInfo"`
	PaymentId           string `json:"paymentId"`
}

var FlypayBStatusCodeMapping = map[int]string{
	100: "authorised",
	200: "decline",
	300: "refunded",
}

func (transaction FlaypayBTransaction) GetAmount() int {
	return transaction.Value
}

func (transaction FlaypayBTransaction) GetCurrency() string {
	return transaction.TransactionCurrency
}

func (transaction FlaypayBTransaction) GetStatus() string {
	return FlypayBStatusCodeMapping[transaction.StatusCode]
}

func (transaction FlaypayBTransaction) GetOrder() string {
	return transaction.OrderInfo
}

func (transaction FlaypayBTransaction) GetPayment() string {
	return transaction.PaymentId
}

func (transaction FlaypayBTransaction) GetProvider() string {
	return "flypayB"
}
