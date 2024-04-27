package main

import "fmt"

func main() {
	fmt.Println(placeOrder("2", 5, 10000))
}

func placeOrder(productID string, qty int, accBalance float64) (bool, float64) {
	if qty > amountInStock(productID) {
		return false, 0.0
	}

	totalPrice := calcPrice(productID, qty)

	if accBalance < totalPrice {
		return false, 0.0
	}

	return true, accBalance - totalPrice
}

func calcPrice(productID string, qty int) float64 {
	return priceList(productID) * float64(qty)
}

func amountInStock(productID string) int {
	if productID == "1" {
		return 11
	}
	if productID == "2" {
		return 25
	}
	if productID == "3" {
		return 4
	}
	if productID == "4" {
		return 6
	}
	if productID == "5" {
		return 50
	}
	if productID == "6" {
		return 2
	}
	if productID == "7" {
		return 0
	}
	if productID == "8" {
		return 99
	}
	if productID == "9" {
		return 1
	}

	return 0
}

func priceList(productID string) float64 {
	if productID == "1" {
		return 1.50
	}
	if productID == "2" {
		return 2.25
	}
	if productID == "3" {
		return 3.00
	}
	if productID == "4" {
		return 1.00
	}
	if productID == "5" {
		return 2.50
	}
	if productID == "6" {
		return 8.99
	}
	if productID == "7" {
		return 22.50
	}
	if productID == "8" {
		return 50.00
	}
	if productID == "9" {
		return 999.99
	}

	return 0.00
}
