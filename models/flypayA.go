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
