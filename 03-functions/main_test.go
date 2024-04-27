package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		tier     string
		expected int
	}

	tests := []testCase{
		{"basic", 10000},
		{"premium", 15000},
		{"enterprise", 50000},
	}

	for _, test := range tests {
		if output := getMonthlyPrice(test.tier); output != test.expected {
			t.Errorf("Test failed: (%s) -> expected: %d actual %d", test.tier, test.expected, output)
		} else {
			fmt.Printf("Test passed: (%s) -> expected: %d actual %d\n", test.tier, test.expected, output)
		}
	}
}
