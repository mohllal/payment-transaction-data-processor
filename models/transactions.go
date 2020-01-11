package models

type Transactions interface {
	GetTrasncations() []Transaction
	Load() Transactions
}
