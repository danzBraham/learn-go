package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	fmt.Println("Interfaces")
	fmt.Println()

	rectangle := Rect{width: 10, height: 15}
	printShapeInfo(rectangle)

	fmt.Println()

	circle := Circle{Radius: 7.3}
	printShapeInfo(circle)

	fmt.Println()
	printNumericValue(6)
	printNumericValue("Zidan")
	printNumericValue(rectangle)
	printNumericValue(circle)
}

/*
Interfaces:
1. Interfaces are implemented implicitly.
If an interface exists and a type has the proper methods defined,
then the type automatically fulfills that interface.
2. A type can implement any number of interfaces (implement multiple interfaces).
*/

// Basic interface
type Default interface{}

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Named paramters interface
type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

// Concrete type Rect implements the Shape interface
type Rect struct {
	width, height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Concrete type Circle implements the Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func printShapeInfo(s Shape) {
	/*
		Type assertion (similar like in Typescript)
		eg. const greetings: string = "Hello"
				greetings.toUpperCase()
		A way to explicitly check and retrieve the concrete value and type of an interface value.
		Used to access the underlying value and type of an interface value when you know the concrete type that the value holds.
	*/

	if c, ok := s.(Circle); ok {
		fmt.Printf("Radius: %.2f\n", c.Radius)
		fmt.Printf("Area: %.2f\n", s.Area())
		fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	}

	if r, ok := s.(Rect); ok {
		fmt.Printf("Width: %.2f, Height: %.2f\n", r.width, r.height)
		fmt.Printf("Area: %.2f\n", s.Area())
		fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	}
}

func printNumericValue(num interface{}) {
	/*
		Type switches
		to do several type assertions in a series.
	*/
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	case Circle:
		fmt.Printf("Radius: %.2f\n", v.Radius)
		fmt.Printf("Area: %.2f\n", v.Area())
		fmt.Printf("Perimeter: %.2f\n", v.Perimeter())
	case Rect:
		fmt.Printf("Width: %.2f, Height: %.2f\n", v.width, v.height)
		fmt.Printf("Area: %.2f\n", v.Area())
		fmt.Printf("Perimeter: %.2f\n", v.Perimeter())
	default:
		fmt.Printf("%T\n", v)
	}
}

/*
Clean Interfaces:
1. Keep interfaces small
Interfaces are meant ot define the minimal behavior necessary to accurately represent an idea or concept.
*/
type File interface {
	io.Closer
	io.Reader
	io.Seeker
	ReadDir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

/*
Any type that satisfies the interface’s behaviors can be considered as a File.
This is convenient because the package doesn’t need to know if it’s dealing with a file on disk,
a network buffer, or a simple []byte.
*/

/*
2. Interfaces should have no knowledge of satisfying types
An interface should define what is necessary for other types to classify as a member of that interface.
They shouldn’t be aware of any types that happen to satisfy the interface at design time.
*/

// not clean interface example:
type NotCleanCar interface {
	Color() string
	Speed() int
	IsFireTruck() bool // this is an anti-pattern
}

// clean interface example:
type CleanCar interface {
	Color() string
	Speed() int
}

type FireTruck interface {
	CleanCar
	IsFireTruck() bool
}

/*
3. Interfaces are not Classes:
	- They are simpler
	- Do not have constructors or deconstructors
	- Are not hierarchical by nature, though there is syntactic sugar
		to create interfaces that happen to be supersets of other interfaces.
	- Interfaces define function signatures, not underlying behavior.
		For example, if five types satisfy the fmt.Stringer interface,
		they all need their own version of the String() function.
*/
