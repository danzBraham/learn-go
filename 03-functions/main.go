package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Functions")

	/*
		Passing variables by value
		"Pass by value" means that when a variable is passed into a function,
		that function receives a copy of the variable.
		The function is unable to mutate the caller's original data.
	*/
	x := 7
	increment(x)
	fmt.Println(x)

	// Ignoring return values
	z, _ := getPoint()
	fmt.Println(z)
}

// Declaring function
func add(x, y int) int {
	return x + y
}

// Multiple parameters
func addToDB(id, name, email string) bool {
	return true
}

func getMonthlyPrice(tier string) int {
	if tier == "basic" {
		return 100 * 100
	} else if tier == "premium" {
		return 150 * 100
	} else if tier == "enterprise" {
		return 500 * 100
	}

	return 0
}

func increment(x int) {
	x++
}

func getPoint() (z time.Time, y int) {
	z = time.Now()
	y = 8
	return z, y
}

/*
Named return values
Good for documentation understanding
you can return explicitly or implicitly (naked return)
you can early returns
*/
func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("can;t divide by zero")
	}

	mul = a * b
	div = a / b
	return mul, div, nil
}
