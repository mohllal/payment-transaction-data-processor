package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FlypayBTransactions struct {
	Transactions []FlypayBTransaction `json:"transactions"`
}

type FlypayBTransaction struct {
	Value               int    `json:"value"`
	TransactionCurrency string `json:"transactionCurrency"`
	StatusCode          int    `json:"statusCode"`
	OrderInfo           string `json:"orderInfo"`
	PaymentId           string `json:"paymentId"`
}

var FlypayBStatusCodeMapping = map[int]string{
	100: "authorised",
	200: "decline",
	300: "refunded",
}

func (transaction FlypayBTransaction) GetAmount() int {
	return transaction.Value
}

func (transaction FlypayBTransaction) GetCurrency() string {
	return transaction.TransactionCurrency
}

func (transaction FlypayBTransaction) GetStatus() string {
	return FlypayBStatusCodeMapping[transaction.StatusCode]
}

func (transaction FlypayBTransaction) GetOrder() string {
	return transaction.OrderInfo
}

func (transaction FlypayBTransaction) GetPayment() string {
	return transaction.PaymentId
}

func (transaction FlypayBTransaction) GetProvider() string {
	return "flypayB"
}

func (transcations FlypayBTransactions) GetTrasncations() []Transaction {
	reuslt := make([]Transaction, len(transcations.Transactions))

	// https://github.com/golang/go/wiki/InterfaceSlic
	// loop through flyB transactions
	for i, transaction := range transcations.Transactions {
		reuslt[i] = transaction
	}

	return reuslt
}

func (transcations FlypayBTransactions) Load() Transactions {
	// open flypayB json file
	jsonFile, err := os.Open("./data/flypayB.json")

	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened flypayB.json")
	// defer the closing operation
	defer jsonFile.Close()

	// read the json file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var flypayB FlypayBTransactions
	// unmarshal the byteArray which contains the json content
	json.Unmarshal(byteValue, &flypayB)

	return flypayB
}
