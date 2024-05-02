package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "what the fubb?! that witch is shiz as fubb"
	removeProfanity(&message)
	fmt.Println(message)
}

func removeProfanity(message *string) {
	if message == nil {
		return
	}

	profanity := map[string]string{
		"fubb":  "****",
		"shiz":  "****",
		"witch": "*****",
	}

	for prf, beep := range profanity {
		*message = strings.ReplaceAll(*message, prf, beep)
	}
}
