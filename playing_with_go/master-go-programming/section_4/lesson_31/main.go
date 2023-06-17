package main

import "fmt"

func main() {
	// ARRAY TYPE

	var numbers = [4]int{4, 5, -9, 100} // array of length 4 with given values
	fmt.Printf("%T\n", numbers)

	// SLICE TYPE

	var cities = []string{"London", "Tokyo", "New York"} // slice has a dynamic length
	fmt.Printf("%T\n", cities)

	// MAP TYPE

	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)

	// STRUCT TYPE

	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)

	// POINTER TYPE

	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	// FUNCTION TYPE

	fmt.Printf("%T\n", f)
}

func f() {

}
