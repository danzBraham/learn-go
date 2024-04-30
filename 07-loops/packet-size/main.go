package main

import "fmt"

func main() {
	fmt.Println(getPacketSize("Hello There"))
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func getPacketSize(message string) int {
	messageLen := len(message)
	if messageLen == 0 {
		return 0
	}

	maxComposite := 0

	for n := messageLen; n >= 4; n-- {
		if !isPrime(n) && messageLen%2 == 0 {
			maxComposite = n
			break
		}
	}

	return maxComposite
}
