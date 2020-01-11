package models

type FlypayATranscations struct {
	Transcations []FlaypayATranscation `json:"transactions"`
}

type FlaypayATranscation struct {
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

func (transcation FlaypayATranscation) GetAmount() int {
	return transcation.Amount
}

func (transcation FlaypayATranscation) GetCurrency() string {
	return transcation.Currency
}

func (transcation FlaypayATranscation) GetStatus() string {
	return FlypayAStatusCodeMapping[transcation.StatusCode]
}

func (transcation FlaypayATranscation) GetOrder() string {
	return transcation.OrderReference
}

func (transcation FlaypayATranscation) GetPayment() string {
	return transcation.TransactionId
}
