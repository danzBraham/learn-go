package main

import "fmt"

func main() {
	fmt.Println(calculateFinalBill(0.75, 5687))
}

func calculateBaseBill(costPerMessage float64, messageSent int) float64 {
	return costPerMessage * float64(messageSent)
}

func calculateDiscount(messagesSent int) float64 {
	if messagesSent > 1000 {
		return 0.1
	} else if messagesSent > 5000 {
		return 0.2
	}

	return 0
}

func calculateFinalBill(costPerMessage float64, numOfMessages int) float64 {
	bill := calculateBaseBill(costPerMessage, numOfMessages)
	discount := calculateDiscount(numOfMessages)
	finalBill := bill - (bill * discount)

	return finalBill
}
