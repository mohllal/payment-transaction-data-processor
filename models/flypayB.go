package models

type FlypayBTranscations struct {
	Transcations []FlaypayBTranscation `json:"transactions"`
}

type FlaypayBTranscation struct {
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

func (transcation FlaypayBTranscation) GetAmount() int {
	return transcation.Value
}

func (transcation FlaypayBTranscation) GetCurrency() string {
	return transcation.TransactionCurrency
}

func (transcation FlaypayBTranscation) GetStatus() string {
	return FlypayBStatusCodeMapping[transcation.StatusCode]
}

func (transcation FlaypayBTranscation) GetOrder() string {
	return transcation.OrderInfo
}

func (transcation FlaypayBTranscation) GetPayment() string {
	return transcation.PaymentId
}
