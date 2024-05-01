package main

import "fmt"

func main() {
	/*
		Maps:
			1. Similar like to Javascript objects
			2. The zero value of a map is nil
			3. Can create a map by using a literal or by using the make() function
			4. Any type can be used as the value in a map, but keys may be of any type that is comparable.
				type that comparable are other than slices, maps, and functions.
			5. Like slices, maps hold references to an underlying data structure.
				If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.
			6. Maps can contain maps, creating a nested structure.
	*/
	ages := map[string]int{
		"Helmi": 2024 - 1999,
		"Zidan": 2024 - 2004,
		"Viky":  2024 - 2006,
	}
	for key, value := range ages {
		fmt.Printf("Key: %s; Value: %d\n", key, value)
	}
	fmt.Println(len(ages))

	// Mutation
	ages["Luffy"] = 19              // Insert
	luffyAge := ages["Luffy"]       // Get
	delete(ages, "Luffy")           // Delete
	luffyExist, ok := ages["Luffy"] // Check if a key exist
	if ok {
		fmt.Println(luffyAge)
		fmt.Println(luffyExist)
	} else {
		fmt.Println("Luffy not exist")
	}

	// Nested map
	nestMap := map[string]map[string]int{
		"Day": {
			"Monday":    1,
			"Tuesday":   2,
			"Wednesday": 3,
		},
		"Month": {
			"January":  1,
			"February": 2,
			"March":    3,
		},
	}
	fmt.Println(nestMap)
}
