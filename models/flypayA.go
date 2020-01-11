package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FlypayATransactions struct {
	Transactions []FlypayATransaction `json:"transactions"`
}

type FlypayATransaction struct {
	Amount         int    `json:"amount"`
	Currency       string `json:"currency"`
	StatusCode     int    `json:"statusCode"`
	OrderReference string `json:"orderReference"`
	TransactionId  string `json:"transactionId"`
}

var FlypayAStatusCodeMapping = map[int]string{
	1: "authorised",
	2: "decline",
	3: "refunded",
}

func (transaction FlypayATransaction) GetAmount() int {
	return transaction.Amount
}

func (transaction FlypayATransaction) GetCurrency() string {
	return transaction.Currency
}

func (transaction FlypayATransaction) GetStatus() string {
	return FlypayAStatusCodeMapping[transaction.StatusCode]
}

func (transaction FlypayATransaction) GetOrder() string {
	return transaction.OrderReference
}

func (transaction FlypayATransaction) GetPayment() string {
	return transaction.TransactionId
}

func (transcation FlypayATransaction) GetProvider() string {
	return "flypayA"
}

func (transcations FlypayATransactions) GetTrasncations() []Transaction {
	reuslt := make([]Transaction, len(transcations.Transactions))

	// loop through flyA transactions
	// https://github.com/golang/go/wiki/InterfaceSlice
	for i, transaction := range transcations.Transactions {
		reuslt[i] = transaction
	}

	return reuslt
}

func (transcations FlypayATransactions) Load() Transactions {
	// open flypayA json file
	jsonFile, err := os.Open("./data/flypayA.json")

	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened flypayA.json")
	// defer the closing operation
	defer jsonFile.Close()

	// read the json file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var flypayA FlypayATransactions
	// unmarshal the byteArray which contains the json content
	json.Unmarshal(byteValue, &flypayA)

	return flypayA
}
