package main

import "fmt"

func main() {
	formatPractice()
	fmt.Println()
	calculateBalance()
}

func formatPractice() {
	fname := "Zidan"
	lname := "Abraham"
	age := 19
	messageRate := 0.5
	isSubscribed := false
	message := "Excellence is an art won by practice and habit"

	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, isSubscribed: %t, Message: %s",
		fname,
		lname,
		age,
		messageRate,
		isSubscribed,
		message,
	)

	fmt.Println(userLog)
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
