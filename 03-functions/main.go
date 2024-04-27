package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Printf("$%d\n", getMonthlyPrice("basic"))

	x, _ := getPoint()
	fmt.Println(x)

	fmt.Println(calculator(10, 5))
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

func getPoint() (x time.Time, y int) {
	x = time.Now()
	y = 8
	return
}

func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("can;t divide by zero")
	}

	mul = a * b
	div = a / b
	return mul, div, nil
}
