package main

import "errors"

type User struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func deleteIfNecessary(users map[string]User, name string) (deleted bool, err error) {
	user, ok := users[name]

	if !ok {
		return false, errors.New("not found")
	}

	if !user.scheduledForDeletion {
		return false, nil
	}

	delete(users, name)
	return true, nil
}
