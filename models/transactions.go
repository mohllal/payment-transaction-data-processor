package models

type TransactionsObject interface {
	GetTrasncations() []Transaction
	Load() TransactionsObject
}
