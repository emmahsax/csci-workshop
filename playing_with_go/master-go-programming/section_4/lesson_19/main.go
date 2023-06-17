package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2

	// Compile error can't assign use a float as an int
	//   a = b
	// The type of a variable can't be changed once declared UNLESS you convert it:
	a = int(b)
	fmt.Println(a, b)

	var x int
	// x = "5" // Compile error because this is a string
	x = 5
	fmt.Println("x:", x)

	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)
}
