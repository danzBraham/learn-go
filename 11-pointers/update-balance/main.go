package main

import "errors"

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type customer struct {
	id      int
	balance float64
}

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(customer *customer, transaction transaction) error {
	if customer.balance < transaction.amount {
		return errors.New("insufficient funds")
	}

	switch transaction.transactionType {
	case transactionDeposit:
		customer.balance += transaction.amount
	case transactionWithdrawal:
		customer.balance -= transaction.amount
	default:
		return errors.New("unknown transaction type")
	}

	return nil
}
