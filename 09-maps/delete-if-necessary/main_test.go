package main

import (
	"testing"
)

func TestDeleteIfNecessary(t *testing.T) {
	type testCase struct {
		users             map[string]User
		name              string
		expectedErrString string
		expectedUsers     map[string]User
		expectedDeleted   bool
	}
	tests := []testCase{
		{
			map[string]User{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			"Erwin",
			"",
			map[string]User{"Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			true,
		},
		{
			map[string]User{"Erwin": {"Erwin", 14355550987, false}, "Levi": {"Levi", 98765550987, false}, "Hanji": {"Hanji", 18265554567, false}},
			"Erwin",
			"",
			map[string]User{"Erwin": {"Erwin", 14355550987, false}, "Levi": {"Levi", 98765550987, false}, "Hanji": {"Hanji", 18265554567, false}},
			false,
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				map[string]User{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
				"Eren",
				"not found",
				map[string]User{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
				false,
			},
		}...)
	}

	for _, test := range tests {
		deleted, err := deleteIfNecessary(test.users, test.name)
		if test.expectedErrString != "" {
			if err == nil {
				t.Errorf("Test Failed: expected error but got none")
			} else if err.Error() != test.expectedErrString {
				t.Errorf("Test Failed: expected error %v but got %v", test.expectedErrString, err)
			}
		} else if err != nil {
			t.Errorf("Test Failed: expected no error but got %v", err)
		} else if !compareMaps(test.users, test.expectedUsers) {
			t.Errorf(
				"Test Failed: \nexpected users: %v \nactual: %v",
				test.expectedUsers,
				test.users,
			)
		} else if deleted != test.expectedDeleted {
			t.Errorf(
				"Test Failed: \nexpected deleted: %v \nactual: %v",
				test.expectedDeleted,
				deleted,
			)
		}
	}
}

func compareMaps(map1, map2 map[string]User) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value1 := range map1 {
		if value2, exist := map2[key]; !exist || value1 != value2 {
			return false
		}
	}
	return true
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
