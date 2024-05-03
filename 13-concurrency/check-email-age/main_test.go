package main

import (
	"fmt"
	"slices"
	"testing"
	"time"
)

func Test(t *testing.T) {
	type testCase struct {
		emails [3]Email
		isOld  [3]bool
	}
	tests := []testCase{
		{[3]Email{
			{
				Body: "Words are pale shadows of forgotten names. As names have power, words have power.",
				Date: time.Date(2019, 2, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				Body: "It's like everyone tells a story about themselves inside their own head. ",
				Date: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				Body: "Bones mend. Regret stays with you forever.",
				Date: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			},
		}, [3]bool{true, false, false}},
		{[3]Email{
			{
				Body: "Music is a proud, temperamental mistress.",
				Date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				Body: "Have you heard of that website Boot.dev?",
				Date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				Body: "It's awesome honestly.",
				Date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
			},
		}, [3]bool{true, true, true}},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{[3]Email{
				{
					Body: "I have stolen princesses back from sleeping barrow kings.",
					Date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
				},
				{
					Body: "I burned down the town of Trebon",
					Date: time.Date(2019, 6, 6, 0, 0, 0, 0, time.UTC),
				},
				{
					Body: "I have spent the night with Felurian and left with both my sanity and my life.",
					Date: time.Date(2022, 7, 0, 0, 0, 0, 0, time.UTC),
				},
			}, [3]bool{true, true, false}},
		}...)
	}

	for _, test := range tests {
		isOld := checkEmailAge(test.emails)
		if !slices.Equal(isOld[:], test.isOld[:]) {
			t.Errorf(`Test Failed:
input:
* %v | %v
* %v | %v
* %v | %v
=>
expected: %v
actual: %v
`,
				test.emails[0].Body, test.emails[0].Date,
				test.emails[1].Body, test.emails[1].Date,
				test.emails[2].Body, test.emails[2].Date,
				test.isOld, isOld,
			)
		} else {
			fmt.Printf(`Test Passed:
input:
* %v | %v
* %v | %v
* %v | %v
=>
expected: %v
actual: %v
`,
				test.emails[0].Body, test.emails[0].Date,
				test.emails[1].Body, test.emails[1].Date,
				test.emails[2].Body, test.emails[2].Date,
				test.isOld, isOld,
			)
		}
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
