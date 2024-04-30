package main

import "fmt"

func main() {
	// Basic loop
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	fmt.Println()

	fmt.Println("max messages:", maxMessages(105))

	/*
		There is no while loop in Go
		Because Go allows for the omission of sections of a for loop,
		a while loop is just a for loop that only has a CONDITION.
	*/
	plantHeight := 0
	for plantHeight < 5 {
		fmt.Printf("The plant is still growing. Current height: %d\n", plantHeight)
		plantHeight++
	}
	fmt.Printf("The plant has grown to: %d\n", plantHeight)

	fmt.Println()

	fizzbuzz(15)

	fmt.Println()

	/*
		Continue
		The continue keyword stops the current iteration of a loop and continues to the next iteration.
		continue is a powerful way to use the guard clause pattern within loops.
	*/
	fmt.Println("Continue:")
	for y := 0; y < 10; y++ {
		if y%2 == 0 {
			continue
		}
		fmt.Println(y)
	}

	fmt.Println()

	/*
		Break
		The break keyword stops the current iteration of a loop and exits the loop.
	*/
	fmt.Println("Break:")
	for z := 0; z < 10; z++ {
		if z == 5 {
			break
		}
		fmt.Println(z)
	}

	fmt.Println()

	printPrimes(20)
}

func maxMessages(thresh int) int {
	totalCost := 0
	/*
		Loops in Go can omit sections of a for loop.
		For example, the CONDITION (middle part) can be omitted which causes the loop to run forever.
	*/
	for i := 0; ; i++ {
		totalCost += 100 + i
		if totalCost > thresh {
			return i
		}
	}
}

func fizzbuzz(num int) {
	// Go supports the standard modulo operator
	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}

		if n%2 == 0 {
			continue
		}

		isPrime := true
		for i := 3; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
			isPrime = true
		}

		if isPrime {
			fmt.Println(n)
		}
	}
}
