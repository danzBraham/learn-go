package main

import (
	"fmt"
)

/*
Structs:
1. Represent structured data
2. Can be nested
3. Similar like object literals in javascript
4. Can be embed (similar like spread operator in javascript)
*/
type User struct {
	username string
	number   int
}

type MessageToSend struct {
	message   string
	sender    User
	recipient User
}

type Car struct {
	brand string
	model string
}

type Truck struct {
	/*
		"car" is embedded, so the definition of a
		"truck" now also additionally contains all
		of the fields of the car struct
	*/
	Car
	bedSize int
}

// Struct methods: function that have a receiver
type Rect struct {
	width  int
	height int
}

func (r Rect) area() int {
	return r.width * r.height
}

type AuthInfo struct {
	username string
	password string
}

func (authI AuthInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", authI.username, authI.password)
}

func main() {
	// Instantiate structs
	myMessage := MessageToSend{
		message: "Hello there!",
		sender: User{
			username: "Zidan",
			number:   123456,
		},
		recipient: User{
			username: "Abraham",
			number:   654321,
		},
	}

	fmt.Println(myMessage)

	/*
		Anonymous structs:
		1. Use it when only meant to be used once
		2. Prefer using named structs as possible
		3. You can nest anonymous struct as field
	*/
	myCar := struct {
		brand string
		model string
		wheel struct {
			radius   int
			material string
		}
	}{
		brand: "Tesla",
		model: "Y3",
		wheel: struct {
			radius   int
			material string
		}{
			radius:   7,
			material: "steel",
		},
	}

	fmt.Println(myCar)

	/*
		Embedded vs Nested:
		1. Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields
		2. Like nested structs, you assign the promoted fields with the embedded struct in a composite literal
	*/
	zidanTruck := Truck{
		bedSize: 17,
		Car: Car{
			brand: "BMW",
			model: "M4",
		},
	}

	fmt.Println(zidanTruck.bedSize)
	fmt.Println(zidanTruck.brand)
	fmt.Println(zidanTruck.model)

	myRect := Rect{
		width:  7,
		height: 10,
	}

	fmt.Println(myRect.area())

	myAuth := AuthInfo{
		username: "danzBraham",
		password: "abc123",
	}

	fmt.Println(myAuth.getBasicAuth())
}

func canSendMessage(mToSend MessageToSend) bool {
	if mToSend.sender.username == "" {
		return false
	}
	if mToSend.recipient.username == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}

	return true
}
