package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name           string
		membershipType membershipType
	}{
		{"Syl", TypeStandard},
		{"Pattern", TypePremium},
		{"Pattern", TypeStandard},
		{"Zidan", TypePremium},
	}
	if withSubmit {
		submitCases := []struct {
			name           string
			membershipType membershipType
		}{
			{"Renarin", TypeStandard},
			{"Lift", TypePremium},
			{"Dhalinar", TypeStandard},
		}
		tests = append(tests, submitCases...)
	}

	for _, tc := range tests {
		user := newUser(tc.name, tc.membershipType)

		msgCharLimit := 100
		if tc.membershipType == TypePremium {
			msgCharLimit = 1000
		}

		if user.Name != tc.name {
			t.Errorf("\nTest Failed: Expected name: %v, got: %s\n",
				tc.name, user.Name)
		}

		if user.Membership.Type != tc.membershipType {
			t.Errorf("\nTest Failed: Expected membership: %v, got: %s\n",
				tc.membershipType, user.Membership.Type)
		}

		if user.Membership.MessageCharLimit != msgCharLimit {
			t.Errorf("\nTest Failed: Expected message character limit: %v, got: %v\n",
				msgCharLimit, user.Membership.MessageCharLimit)
		}

		fmt.Printf("\nTest Passed: user: %s, membership type: %s, message character limit: %v\n",
			user.Name, user.Membership.Type, user.Membership.MessageCharLimit)
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
