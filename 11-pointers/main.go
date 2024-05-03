package main

import "fmt"

func main() {
	/*
		Pointer
		Hold the memory address of a value.
		A variable that stores the memory address of another variable.
		This means that a pointer "points to" the location of where the data is stored NOT the actual data itself.
		A pointer's zero value is nil.

		Why are pointers useful?
		Pointers allow us to manipulate data in memory directly, without making copies or duplicating data.
		This can make programs more efficient and allow us to do things that would be difficult or impossible without them.

		JUST BECAUSE YOU CAN DOESN'T MEAN YOU SHOULD!!
		Because pointers can be very dangerous.
		Whenever you're dealing with pointers you should check if it's nil before trying to dereference it.
		It's generally a better idea to have your functions accept non-pointers
		and return new values rather than mutating pointer inputs.
	*/

	// The * syntax defines a pointer:
	// var p *int

	// The & operator generates a pointer to its operand.
	myString := "Hello"
	myStringPtr := &myString // store the address of myString
	fmt.Println(myStringPtr)
	// The * dereferences a pointer to gain access to the value
	fmt.Println(*myStringPtr)

	lambo := Car{0, "Orange"}
	fmt.Println(lambo)
	lambo.setSpeed(345)
	fmt.Println(lambo)

	address1 := Addres{"Denpasar", "Bali", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Singaraja"
	address3 = &Addres{"Denpasar", "Bali", "Indonesia"}
	*address2 = Addres{"Bedugul", "Bali", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}

type Car struct {
	Speed int
	Color string
}

/*
Pointer receiver
A receiver type on a method can be a pointer.
Methods with pointer receivers can modify the value to which the receiver points.
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
*/
func (c *Car) setSpeed(spd int) {
	c.Speed = spd
}

type Addres struct {
	City, Province, Country string
}
