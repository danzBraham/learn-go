package main

import "fmt"

func main() {
	var startup string = "Textio SMS service booting up..."
	var message string = "Sending text message"
	var confirm string = "Mesage sent!"

	fmt.Println(startup)
	fmt.Println(message)
	fmt.Println(confirm)

	sender := "Zidan"
	recipient := "Abraham"
	chat := "Hello, I am Zidan"

	fmt.Printf("%s to %s: %s\n", sender, recipient, chat)
}
