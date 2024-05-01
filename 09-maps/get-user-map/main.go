package main

import "errors"

type User struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]User, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	users := make(map[string]User, len(names))

	for i, name := range names {
		users[name] = User{name: name, phoneNumber: phoneNumbers[i]}
	}

	return users, nil
}
