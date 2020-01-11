package utils

import "github.com/Mohllal/go-challange/models"

var GenerateFakeTransactionsByCriteria = func(criteria string) []models.Transaction {
	transactions := make([]models.Transaction, 0)

	if criteria == "provider" {
		transactions = generateFakeTransactionsWithTheSameProvider()
	}

	return transactions
}

// generate fake transactions with the same provider
var generateFakeTransactionsWithTheSameProvider = func() []models.Transaction {
	fake1 := generateFakeTranscation(100, "EGP", "authorised", "2e58bd43-7abb", "4dc2-a8a1")
	fake2 := generateFakeTranscation(500, "US", "decline", "2e58uim3-7op", "4dc2-ui9o")

	transactions := []models.Transaction{fake1, fake2}
	return transactions
}

// generate fake transaction
var generateFakeTranscation = func(amount int, currency string, status string, order string, transaction string) models.Transaction {
	transactionMock := TransactionFake{
		Amount:      amount,
		Currency:    currency,
		Status:      status,
		Order:       order,
		Transaction: transaction,
	}

	return transactionMock
}
