package main

import "fmt"

func main() {
	formatPractice()
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
