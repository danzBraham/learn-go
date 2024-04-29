package main

import (
	// The Go standard library provides an "errors" package that makes it easy to deal with errors.
	"errors"
	"fmt"
	"strconv"
)

func main() {
	/*
		When something can go wrong in a function,
		that function should return an error as its last return value.
		Any code that calls a function that can return an error should handle errors by testing whether the error is nil.
	*/
	i, err := strconv.Atoi("2983")
	if err != nil {
		fmt.Println("couldn't convert:", err)
		return
	}

	fmt.Println(i)

	// Go std lib for dealing with errors
	var errs error = errors.New("something went wrong")
	fmt.Println(errs)
}

/*
An Error is any type that implements the simple built-in error interface.
That is why you can build your own custom types that implement the error interface.
Here's an example of a userError struct that implements the error interface:
*/
type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%s has a problem with their account", e.name)
}

func sendSMS(msg, userName string) (string, error) {
	if msg == "" {
		return "", userError{name: userName}
	}

	return msg, nil
}
