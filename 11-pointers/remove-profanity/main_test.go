package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		messageIn *string
		expected  *string
	}
	s1 := "English, motherfubber, do you speak it?"
	s2 := "English, mother****er, do you speak it?"
	s3 := "Does he look like a witch?"
	s4 := "Does he look like a *****?"

	tests := []testCase{
		{
			&s1,
			&s2,
		},
		{
			nil,
			nil,
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				&s3,
				&s4,
			},
			{
				nil,
				nil,
			},
		}...)
	}

	for _, test := range tests {
		var original *string
		if test.messageIn != nil {
			originalVal := *test.messageIn
			original = &originalVal
		}
		removeProfanity(test.messageIn)
		if test.messageIn != nil &&
			test.expected != nil &&
			original != nil &&
			*test.messageIn != *test.expected {
			t.Errorf(`Test Failed:
input:
%v
=>
expected:
%v
actual:
%v
===========================
`,
				*original,
				*test.expected,
				*test.messageIn,
			)
		} else if (test.messageIn == nil || test.expected == nil) &&
			test.messageIn != test.expected {
			t.Errorf(`Test Failed:
input:
%v
=>
expected:
%v
actual:
%v
===========================
`,
				original,
				test.expected,
				test.messageIn,
			)
		} else if test.messageIn == nil && test.expected == nil {
			fmt.Printf(`Test Passed:
input:
%v
=>
expected:
%v
actual:
%v
===========================
`,
				original,
				test.expected,
				test.messageIn,
			)
		} else {
			fmt.Printf(`Test Passed:
  input:
%v
=>
  expected:
%v
  actual:
%v
===========================
`,
				*original,
				*test.expected,
				*test.messageIn,
			)
		}
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
