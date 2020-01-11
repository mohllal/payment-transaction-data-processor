package data

import (
	"code-challange/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var LoadFlyATranscations = func() models.FlypayATranscations {
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

	var flypayA models.FlypayATranscations
	// unmarshal the byteArray which contains the json content
	json.Unmarshal(byteValue, &flypayA)

	return flypayA
}

var LoadFlyBTranscations = func() models.FlypayBTranscations {
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

	var flypayB models.FlypayBTranscations
	// unmarshal the byteArray which contains the json content
	json.Unmarshal(byteValue, &flypayB)

	return flypayB
}

var LoadAllTranscations = func () []models.TranscationObject {
	var flyATranscations = LoadFlyATranscations()
	var flyBTranscations = LoadFlyBTranscations()

	// https://github.com/golang/go/wiki/InterfaceSlice
	var transcations []models.TranscationObject = make([]models.TranscationObject, len(flyATranscations.Transcations)+len(flyBTranscations.Transcations))

	// loop throug flyA transcations
	for i, transcation := range flyATranscations.Transcations {
		transcations[i] = transcation
	}

	// loop throug flyB transcations
	for i, transcation := range flyBTranscations.Transcations {
		transcations[i+len(flyATranscations.Transcations)] = transcation
	}

	return transcations;
}