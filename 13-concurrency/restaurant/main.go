package main

import (
	"fmt"
	"time"
)

func main() {
	// Create new order
	order := Order{
		ID:         1,
		Appetizer:  "Bruschetta",
		MainCourse: "Pasta Carbonara",
		Dessert:    "Tiramisu",
	}

	// Create goroutines to handle different parts of the order concurrently
	go prepareAppetizer(order)
	go prepareMainCourse(order)
	go prepareDessert(order)

	// Wait for all goroutines to finish
	for i := 0; i < 3; i++ {
		completed := <-Kitchen
		fmt.Printf("Order #%d is completed\n", completed.ID)
	}
}

// Order represents a customer's order
type Order struct {
	ID         int
	Appetizer  string
	MainCourse string
	Dessert    string
}

// Kitchen is a channel to send and receive orders
var Kitchen = make(chan Order)

func prepareAppetizer(order Order) {
	time.Sleep(1 * time.Second) // Simulate preparation time
	fmt.Printf("Appetizer '%s' for Order #%d is ready!\n", order.Appetizer, order.ID)
	Kitchen <- order // Send the order back to the kitchen
}

func prepareMainCourse(order Order) {
	time.Sleep(2 * time.Second) // Simulate preparation time
	fmt.Printf("Main Course '%s' for Order #%d is ready!\n", order.Appetizer, order.ID)
	Kitchen <- order // Send the order back to the kitchen
}

func prepareDessert(order Order) {
	time.Sleep(1 * time.Second) // Simulate preparation time
	fmt.Printf("Dessert '%s' for Order #%d is ready!\n", order.Appetizer, order.ID)
	Kitchen <- order // Send the order back to the kitchen
}
