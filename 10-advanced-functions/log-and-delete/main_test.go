package main

import (
	"fmt"
	"testing"
)

func TestLogAndDelete(t *testing.T) {
	type testCase struct {
		users         map[string]User
		name          string
		expectedUsers map[string]User
		expectedLog   string
	}
	tests := []testCase{
		{
			users: map[string]User{
				"laura": {name: "laura", number: 4355556023, admin: false},
				"dale":  {name: "dale", number: 8015558937, admin: true},
			},
			name: "laura",
			expectedUsers: map[string]User{
				"dale": {name: "dale", number: 8015558937, admin: true},
			},
			expectedLog: logDeleted,
		},
		{
			users: map[string]User{
				"audrey": {name: "audrey", number: 4355556024, admin: false},
				"bob":    {name: "bob", number: 8015558938, admin: true},
				"bobby":  {name: "bobby", number: 8015558939, admin: true},
			},
			name: "bobby",
			expectedUsers: map[string]User{
				"audrey": {name: "audrey", number: 4355556024, admin: false},
				"bob":    {name: "bob", number: 8015558938, admin: true},
			},
			expectedLog: logAdmin,
		},
	}

	if withSubmit {
		tests = append(tests, testCase{
			users: map[string]User{
				"log lady": {name: "log lady", number: 4355556025, admin: false},
				"shelly":   {name: "shelly", number: 8015558940, admin: false},
				"leland":   {name: "leland", number: 8015558941, admin: false},
			},
			name: "laura",
			expectedUsers: map[string]User{
				"log lady": {name: "log lady", number: 4355556025, admin: false},
				"shelly":   {name: "shelly", number: 8015558940, admin: false},
				"leland":   {name: "leland", number: 8015558941, admin: false},
			},
			expectedLog: logNotFound,
		})
	}

	for _, test := range tests {
		originalMap := getMapCopy(test.users)
		log := logAndDelete(test.users, test.name)
		if log != test.expectedLog ||
			!compareMaps(test.users, test.expectedUsers) {
			t.Errorf(`
Test Failed.
  users:
%v
  name: %v
	->
  expectedUsers:
%v
  actualUsers:
%v
  expectedLog: %s
  actualLog: %s
`,
				formatMap(originalMap),
				test.name,
				formatMap(test.expectedUsers),
				formatMap(test.users),
				test.expectedLog,
				log,
			)
		} else {
			fmt.Printf(`
Test Passed.
  users:
%v
  name: %v
  ->
  expectedUsers:
%v
  actualUsers:
%v
  expectedLog: %s
  actualLog: %s
`,
				formatMap(originalMap),
				test.name,
				formatMap(test.expectedUsers),
				formatMap(test.users),
				test.expectedLog,
				log,
			)
		}
	}
}

func getMapCopy(m map[string]User) map[string]User {
	copy := make(map[string]User)
	for key, value := range m {
		copy[key] = value
	}
	return copy
}

func formatMap(m map[string]User) string {
	var str string
	for key, value := range m {
		str += fmt.Sprintf("  * %s: %v\n", key, value)
	}
	return str
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
