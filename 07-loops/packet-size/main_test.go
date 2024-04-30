package main

import (
	"fmt"
	"testing"
)

func TestGetPacketSize(t *testing.T) {
	type testCase struct {
		message            string
		expectedPacketSize int
	}
	testCases := []testCase{
		{"No, this is Boots the Magic Bear", 32},
		{"I've been trying to reach you about your car's extended warranty", 64},
	}
	if withSubmit {
		testCases = append(testCases,
			[]testCase{
				{"", 0},
				{"Hello there", 0},
				{"I've got a bad conditional statement about this.", 48},
				{"It's over, Ballan, I have the high ground. Don't try it!", 56},
			}...,
		)
	}

	for _, tc := range testCases {
		result := getPacketSize(tc.message)
		if result != tc.expectedPacketSize {
			t.Errorf(`Test Failed: Message: "%s"
=>
* expected: %v
* actual: %v
`,
				tc.message,
				tc.expectedPacketSize,
				result,
			)
		} else {
			fmt.Printf(`Test Passed: Message: "%s"
=>
* expected: %v
* actual: %v
`,
				tc.message,
				tc.expectedPacketSize,
				result,
			)
		}
	}
}

var withSubmit = true
