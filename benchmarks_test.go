package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mohllal/go-challenge/data"
	"github.com/gin-gonic/gin"
)

func BenchmarkTransactionRoute(b *testing.B) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()
	gin.DefaultWriter = ioutil.Discard

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/payment/transaction", nil)
		router.ServeHTTP(w, req)
	}
}

func BenchmarkTransactionsFilterWithProvider(b *testing.B) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()
	gin.DefaultWriter = ioutil.Discard

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/payment/transaction?provider=flypayA", nil)
		router.ServeHTTP(w, req)
	}
}

func BenchmarkTransactionsFilterWithStatus(b *testing.B) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()
	gin.DefaultWriter = ioutil.Discard

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/payment/transaction?statusCode=authorised", nil)
		router.ServeHTTP(w, req)
	}
}

func BenchmarkTransactionsFilterWithCurrency(b *testing.B) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()
	gin.DefaultWriter = ioutil.Discard

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/payment/transaction?currency=AUD", nil)
		router.ServeHTTP(w, req)
	}
}

func BenchmarkTransactionsFilterWithinAmount(b *testing.B) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()
	gin.DefaultWriter = ioutil.Discard

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/payment/transaction?amountMin=100&amountMax=200", nil)
		router.ServeHTTP(w, req)
	}
}
func BenchmarkLoadAllTransactions(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	for i := 0; i < b.N; i++ {
		data.LoadAllTransactions()
	}
}
