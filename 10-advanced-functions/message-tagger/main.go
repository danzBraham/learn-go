package main

import (
	"strings"
)

// func main() {
// 	fmt.Println(strings.Contains("HELLO", "HELLO"))
// }

type SMS struct {
	id      string
	content string
	tags    []string
}

func tagger(msg SMS) []string {
	tags := []string{}
	lowerMsg := strings.ToLower(msg.content)
	substrings := map[string]string{
		"urgent": "Urgent",
		"sale":   "Promo",
	}

	for sub, tag := range substrings {
		if strings.Contains(lowerMsg, sub) {
			tags = append(tags, tag)
		}
	}

	return tags
}

// Before refactor (pros: more readable; cons: not optimized)
func tagMessages(messages []SMS, tagger func(SMS) []string) []SMS {
	result := []SMS{}
	for _, msg := range messages {
		msg.tags = tagger(msg)
		result = append(result, msg)
	}
	return result
}

// After refactor (pros: more optimized; cons: not readable)
// func tagMessages(messages []SMS, tagger func(SMS) []string) []SMS {
// 	result := make([]SMS, len(messages))
// 	for i, msg := range messages {
// 		result[i] = SMS{
// 			id:      msg.id,
// 			content: msg.content,
// 			tags:    tagger(msg),
// 		}
// 	}
// 	return result
// }
