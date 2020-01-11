package data

import (
	"code-challange/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var LoadFlyATransactions = func() models.FlypayATransactions {
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

	var flypayA models.FlypayATransactions
	// unmarshal the byteArray which contains the json content
	json.Unmarshal(byteValue, &flypayA)

	return flypayA
}

var LoadFlyBTransactions = func() models.FlypayBTransactions {
	// open flypayA json file
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

	var flypayB models.FlypayBTransactions
	// unmarshal the byteArray which contains the json content
	json.Unmarshal(byteValue, &flypayB)

	return flypayB
}

var LoadAllTransactions = func () []models.TransactionObject {
	var flyATransactions = LoadFlyATransactions()
	var flyBTransactions = LoadFlyBTransactions()

	// https://github.com/golang/go/wiki/InterfaceSlice
	var transactions []models.TransactionObject = make([]models.TransactionObject, len(flyATransactions.Transactions)+len(flyBTransactions.Transactions))

	// loop through flyA transactions
	for i, transaction := range flyATransactions.Transactions {
		transactions[i] = transaction
	}

	// loop through flyB transactions
	for i, transaction := range flyBTransactions.Transactions {
		transactions[i+len(flyATransactions.Transactions)] = transaction
	}

	// fmt.Println(transactions[0].GetStatus())

	return transactions;
}