package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTagMessages(t *testing.T) {
	tests := []struct {
		messages []SMS
		expected [][]string
	}{
		{
			messages: []SMS{{id: "001", content: "This is urgent please respond!"}, {id: "002", content: "Big sale on all items!"}},
			expected: [][]string{{"Urgent"}, {"Promo"}},
		},
		{
			messages: []SMS{{id: "003", content: "Enjoy your day"}},
			expected: [][]string{{}},
		},
	}

	if withSubmit {
		tests = append(tests, struct {
			messages []SMS
			expected [][]string
		}{
			messages: []SMS{{id: "004", content: "Urgent sale! Don't miss out on these urgent promotions!"}},
			expected: [][]string{{"Urgent", "Promo"}},
		})
	}

	for _, test := range tests {
		actual := tagMessages(test.messages, tagger)
		for i, msg := range actual {
			if !reflect.DeepEqual(msg.tags, test.expected[i]) {
				t.Errorf("Test Failed for message ID %s.\n Expected tags: %v\n Actual tags: %v\n", msg.id, test.expected[i], msg.tags)
			} else {
				fmt.Printf("Test Passed for message ID %s.\n Expected tags: %v\n Actual tags: %v\n", msg.id, test.expected[i], msg.tags)
			}
		}
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
