package utils

type TransactionFake struct {
	Amount      int    `json:"Amount"`
	Currency    string `json:"Currency"`
	Status      string `json:"Status"`
	Order       string `json:"Order"`
	Transaction string `json:"Transaction"`
}

func (transaction TransactionFake) GetAmount() int {
	return transaction.Amount
}

func (transaction TransactionFake) GetCurrency() string {
	return transaction.Currency
}

func (transaction TransactionFake) GetStatus() string {
	return transaction.Status
}

func (transaction TransactionFake) GetOrder() string {
	return transaction.Order
}

func (transaction TransactionFake) GetPayment() string {
	return transaction.Transaction
}

func (transcation TransactionFake) GetProvider() string {
	return "flypayFake"
}
