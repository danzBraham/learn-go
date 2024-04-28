package main

import "fmt"

func main() {
	calculateBalance()
}

func calculateBalance() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	finalCost = bulkMessageCost

	if isPremiumUser {
		finalCost -= finalCost * discountRate
	}

	if accountBalance >= finalCost {
		accountBalance -= finalCost
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}

	fmt.Printf("Account Balance: $%.1f\n", accountBalance)
}
