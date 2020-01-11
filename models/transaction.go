package models

type Transaction interface {
	GetAmount() int
	GetCurrency() string
	GetStatus() string
	GetOrder() string
	GetPayment() string
	GetProvider() string
}
