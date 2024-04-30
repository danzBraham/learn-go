package main

import "unicode"

func isValidPassword(password string) bool {
	passLen := len(password)

	if passLen < 5 || passLen > 12 {
		return false
	}

	hasUpperCase := false
	isContainNumber := false

	for _, p := range password {
		if unicode.IsUpper(p) {
			hasUpperCase = true
			break
		}
	}

	for _, p := range password {
		if unicode.IsNumber(p) {
			isContainNumber = true
			break
		}
	}

	if !hasUpperCase || !isContainNumber {
		return false
	}

	return true
}
