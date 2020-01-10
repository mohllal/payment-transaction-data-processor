package models

type FlypayBTranscations struct {
	Transcations []FlaypayBTranscation `json:"transactions"`
}

type FlaypayBTranscation struct {
	Value               int    `json:"value"`
	TransactionCurrency string `json:"transactionCurrency"`
	StatusCode          int    `json:"statusCode"`
	OrderInfo 			string `json:"orderInfo"`
	PaymentId  			string `json:"paymentId"`
}
