package main

import "fmt"

func main() {
	fmt.Println("add:", aggregate(2, 5, 7, add))
	fmt.Println("mul:", aggregate(2, 5, 7, mul))

	squareFunc := selfMath(mul)
	fmt.Println("squareFunc:", squareFunc(5))
	fmt.Println()

	deff()
	fmt.Println()

	zidanAggregator := concatter()
	zidanAggregator("Hello")
	zidanAggregator("I am")
	fmt.Println(zidanAggregator("Zidan"))

	/*
		Anonymous Function
		A function without name.
		Useful when defining a function that will only be used once or to create a quick closure.
	*/
	nums := []int{2, 3, 5, 7, 11}
	allNumsDoubles := doMath(func(x int) int {
		return x * 2
	}, nums)
	fmt.Println(allNumsDoubles)
}

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

/*
First-Class Functions
When functions in that language are treated like any other variable.
For example, in such a language, a function can be passed as an argument to other functions,
can be returned by another function and can be assigned as a value to a variable.

Higher-Order Functions
A function that returns a function or accepts a function as input is called a Higher-Order Function.

Go supports first-class and higher-order functions.
Another way to think of this is that a function is just another type -- just like ints and strings and bools.
*/
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

/*
Currying
It allows a function with multiple arguments to be transformed into a sequence of functions, each taking a single argument.
Although Go does not support full currying, it is possible to simulate this behavior.
By designing functions that take other functions as inputs and return new functions, we can achieve a similar effect.
*/
func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

/*
Defer
The defer keyword is a fairly unique feature of Go.
It allows a function to be executed automatically just before its enclosing function returns.
Deferred functions are typically used to close database connections, file handlers and the like.
*/
func deff() {
	defer fmt.Println("Finish")
	fmt.Println("Start")
	fmt.Println("Hello")
	fmt.Println("There")
}

/*
Closures
A function that references variables from outside its own function body.
The function may access and assign to the referenced variables.
Basically closures can mutate a variable outside its body
*/
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func doMath(f func(int) int, nums []int) []int {
	result := make([]int, 0)
	for _, num := range nums {
		result = append(result, f(num))
	}
	return result
}
