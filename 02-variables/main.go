package main

import "fmt"

// Declaring variable
var number int
var desimal float64 = 3.14

func main() {
	// Short variable declaration
	name := "Zidan"
	age := 19
	fmt.Println("Short variable declaration:", name, age)

	// Type inference
	x := 7            // int
	y := 8.97         // float64
	z := 0.867 + 0.5i // complex128
	fmt.Println("Type inference:", x, y, z)

	// Same line declaration
	footballClub, goal := "Real Madrid", 5
	fmt.Printf("Same line declaration: %s has scored %d goals\n", footballClub, goal)

	// Type sizez
	// Size of the type indicating bits
	// int  int8  int16  int32  int64 whole numbers
	// uint uint8 uint16 uint32 uint64 uintptr positive whole numbers
	// float32 float64 decimal numbers
	// complex64 complex128 imaginary numbers (rare)

	// PREFER "DEFAULT" TYPES when choosing type unless you concerned about performance and memory usage
	// bool string int uint byte rune float64 complex128
	var accountAge int = 7
	fmt.Println("Type sizez:", accountAge)

	// Constants
	// Must be known at compile time
	const pi = 3.14
	fmt.Println("Constant:", pi)

	// Formatting strings in go
	// default 	= %v
	// string 	= %s
	// int 			= %d
	// float 		= %f
	// bool 		= %t
	s := fmt.Sprintf("Hello my name is %s, and i am %d years old", name, age)
	fmt.Println("Formatting strings:", s)

	// Conditionals
	fmt.Printf("Conditionals: ")
	height := 7
	if height > 6 {
		fmt.Println("You're so tall!")
	} else if height > 4 {
		fmt.Println("You're average tall!")
	} else {
		fmt.Println("You're short!")
	}

	// Initial statement of an if block
	fmt.Printf("Initial statement: ")
	if length := 15; length > 10 {
		fmt.Println("You're message is too long!")
	} else {
		fmt.Println("You're message is valid!")
	}
}
