package main

import "fmt"

func main() {
	/*
		Arrays are fixed-size groups of variables of the same type.
		Once you make an array like [10]int you can't add an 11th element.
		The zero value of array is zero.
	*/
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	/*
		Slices:
		1. Slices is a dynamically-sized, flexible view of the elements of an array.
		2. The zero value of slice is nil.
		3. Slices always have an underlying array, though it isn't always specified explicitly.
		4. Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.

		To explicitly create a slice on top of an array we can do:
			arrayname[lowIndex:highIndex]		return value from lowIndex to highIndex
			arrayname[lowIndex:]						return value from lowIndex to last index
			arrayname[:highIndex]						return value from first index to highIndex
			arrayname[:]										return value of all index
		Where lowIndex is inclusive and highIndex is exclusive.

		To creata basic slice just do: []int
	*/
	mySlice := primes[:]
	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println(mySlice)

	/*
		make function:
		Most of the time we don't need to think about the underlying array of a slice.
		We can create a new slice using the make function.
		Slices created with make will be filled with the zero value of the type.
	*/
	// func make([]T, len, cap int) []T
	makeSlice := make([]int, 5, 10)
	fmt.Println(makeSlice)

	// the capacity argument is usually omitted and defaults to the length
	makeSlice2 := make([]int, 5)
	fmt.Println(makeSlice2)

	// If we want to create a slice with a specific set of values, we can use a slice literal:
	mySlice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice2)

	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))

	fmt.Println()

	fmt.Println("Variadic:", sum(1, 2, 3, 4, 5))

	fmt.Println()

	myStrings := []string{"Hello", "I'm", "Zidan"}
	// Spread operator is the same like in Javascript but the "..." in the end
	printStrings(myStrings...)

	fmt.Println()

	/*
		Slice of Slices
		Slices can hold other slices, effectively creating a matrix, or a 2D slice.
	*/
	rows := [][]int{
		{0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5},
	}
	fmt.Println("rows:", rows)
	fmt.Println("len:", len(rows))
	fmt.Println("cap:", cap(rows))

	fmt.Println()

	/*
		Tricky Slices
		The append() function changes the underlying array of its parameter AND returns a new slice.
		This means that using append() on anything other than itself is usually a BAD idea.

		DON'T DO THIS!
		someSlice = append(otherSlice, element)
		To avoid bugs, always use the append function on the same slice the result is assigned to:
	*/
	sliceKu := []int{}
	sliceKu = append(sliceKu, 1)

	/*
		Range
		for INDEX, VALUE := range someSlice {}
	*/
	for _, p := range primes {
		fmt.Printf("%d,", p)
	}

	fmt.Println()
}

func updateSlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

// Variadic similar like rest operator in Javascript
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func printStrings(str ...string) {
	for _, s := range str {
		fmt.Println(s)
	}
}
