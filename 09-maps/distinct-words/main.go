package main

import (
	"fmt"
	"strings"
)

func main() {
	messages := []string{"Hello world", "hello"}
	fmt.Println(countDistinctWords(messages))
}

func countDistinctWords(messages []string) int {
	msgMap := make(map[string]bool)

	for _, message := range messages {
		msg := strings.Fields(message)
		for _, m := range msg {
			lowerM := strings.ToLower(m)
			if _, ok := msgMap[lowerM]; !ok {
				msgMap[lowerM] = true
			}
		}
	}

	return len(msgMap)
}
