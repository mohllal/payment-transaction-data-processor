package main

import (
	"encoding/json"
	"github.com/Mohllal/go-challenge/data"
	"github.com/Mohllal/go-challenge/models"
	"github.com/Mohllal/go-challenge/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransactionRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()

	// stub the LoadAllTransactions function
	data.LoadAllTransactions = func() []models.Transaction {
		return make([]models.Transaction, 0)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/payment/transaction", nil)
	router.ServeHTTP(w, req)

	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	response := utils.TransactionsResponse{}

	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// grab the value & whether or not it exists
	value := response.Transactions

	// make some assertions on the correctness of the response
	assert.Nil(t, err)
	assert.Equal(t, make([]utils.TransactionFake, 0), value)
}

func TestTransactionsFilterWithProvider(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()

	// get the fake data from the factory
	transactions := utils.GenerateFakeTransactionsByCriteria("provider")

	// stub the LoadAllTransactions function
	data.LoadAllTransactions = func() []models.Transaction {
		return transactions
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/payment/transaction?provider=flypayFake", nil)
	router.ServeHTTP(w, req)

	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	response := utils.TransactionsResponse{}

	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// grab the value & whether or not it exists
	value := response.Transactions

	valueAsTransactionType := make([]models.Transaction, len(value))

	// https://github.com/golang/go/wiki/InterfaceSlic
	// transform back the array
	for i, transaction := range value {
		valueAsTransactionType[i] = transaction
	}

	// make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.Equal(t, transactions, valueAsTransactionType)
}

func TestTransactionsFilterWithStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()

	// get the fake data from the factory
	transactions := utils.GenerateFakeTransactionsByCriteria("status")

	// stub the LoadAllTransactions function
	data.LoadAllTransactions = func() []models.Transaction {
		return transactions
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/payment/transaction?statusCode=authorised", nil)
	router.ServeHTTP(w, req)

	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	response := utils.TransactionsResponse{}

	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// grab the value & whether or not it exists
	value := response.Transactions

	valueAsTransactionType := make([]models.Transaction, len(value))

	// https://github.com/golang/go/wiki/InterfaceSlic
	// transform back the array
	for i, transaction := range value {
		valueAsTransactionType[i] = transaction
	}

	// make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.Equal(t, transactions, valueAsTransactionType)
}

func TestTransactionsFilterWithCurrency(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()

	// get the fake data from the factory
	transactions := utils.GenerateFakeTransactionsByCriteria("currency")

	// stub the LoadAllTransactions function
	data.LoadAllTransactions = func() []models.Transaction {
		return transactions
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/payment/transaction?currency=EGP", nil)
	router.ServeHTTP(w, req)

	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	response := utils.TransactionsResponse{}

	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// grab the value & whether or not it exists
	value := response.Transactions

	valueAsTransactionType := make([]models.Transaction, len(value))

	// https://github.com/golang/go/wiki/InterfaceSlic
	// transform back the array
	for i, transaction := range value {
		valueAsTransactionType[i] = transaction
	}

	// make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.Equal(t, transactions, valueAsTransactionType)
}

func TestTransactionsFilterWithinAmount(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()

	// get the fake data from the factory
	transactions := utils.GenerateFakeTransactionsByCriteria("amount")

	// stub the LoadAllTransactions function
	data.LoadAllTransactions = func() []models.Transaction {
		return transactions
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/payment/transaction?amountMin=100&amountMax=200", nil)
	router.ServeHTTP(w, req)

	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	response := utils.TransactionsResponse{}

	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// grab the value & whether or not it exists
	value := response.Transactions

	valueAsTransactionType := make([]models.Transaction, len(value))

	// https://github.com/golang/go/wiki/InterfaceSlic
	// transform back the array
	for i, transaction := range value {
		valueAsTransactionType[i] = transaction
	}

	// make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.Equal(t, transactions, valueAsTransactionType)
}
