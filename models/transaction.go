package models

type TranscationObject interface {
	GetAmount() int
	GetCurrency() string
	GetStatus() string
	GetOrder() string
	GetPayment() string
}
