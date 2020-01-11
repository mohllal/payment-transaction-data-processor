package data

import (
	"github.com/Mohllal/go-challange/models"
)

var LoadAllTransactions = func() []models.Transaction {
	var flyATransactions = models.Transactions.Load(models.FlypayATransactions{})
	var flyBTransactions = models.Transactions.Load(models.FlypayBTransactions{})

	var transactions []models.Transaction = append(flyATransactions.GetTrasncations(), flyBTransactions.GetTrasncations()...)

	return transactions
}

var FindTransactionsByProvider = func(transactions []models.Transaction, provider string) []models.Transaction {
	transactionsByProvider := make([]models.Transaction, 0)

	// loop through all transactions and return only the ones with the sent provider
	for _, transaction := range transactions {
		if transaction.GetProvider() == provider {
			transactionsByProvider = append(transactionsByProvider, transaction)
		}
	}

	return transactionsByProvider
}

var FindTransactionsByStatusCode = func(transactions []models.Transaction, statusCode string) []models.Transaction {
	transactionsByStatusCode := make([]models.Transaction, 0)

	// loop through all transactions and return only the ones with the sent statusCode
	for _, transaction := range transactions {
		if transaction.GetStatus() == statusCode {
			transactionsByStatusCode = append(transactionsByStatusCode, transaction)
		}
	}

	return transactionsByStatusCode
}

var FindTransactionsByAmount = func(transactions []models.Transaction, amountMin int, amountMax int) []models.Transaction {
	transactionsByAmount := make([]models.Transaction, 0)

	// loop through all transactions and return only the ones between the sent amountMin and amountMax inclusively
	for _, transaction := range transactions {
		if transaction.GetAmount() >= amountMin && transaction.GetAmount() <= amountMax {
			transactionsByAmount = append(transactionsByAmount, transaction)
		}
	}

	return transactionsByAmount
}

var FindTransactionsByCurrency = func(transactions []models.Transaction, currency string) []models.Transaction {
	transactionsByCurrency := make([]models.Transaction, 0)

	// loop through all transactions and return only the ones with the sent currency
	for _, transaction := range transactions {
		if transaction.GetCurrency() == currency {
			transactionsByCurrency = append(transactionsByCurrency, transaction)
		}
	}

	return transactionsByCurrency
}
