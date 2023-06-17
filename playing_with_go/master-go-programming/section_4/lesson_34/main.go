package main

import "fmt"

func main() {
	// COMPARISON OPERATORS

	a, b := 5, 10
	fmt.Println(a == b)  // -> false
	fmt.Println(a != b)  // -> true
	fmt.Println(a > 5)   // -> false
	fmt.Println(a >= 5)  // -> true
	fmt.Println(10 <= a) // -> false

	// LOGICAL OPERATORS

	fmt.Println(a > 1 && b == 10)        // -> true
	fmt.Println(a < 1 && b == 10)        // -> false
	fmt.Println(a == 5 || b == 100)      // -> true
	fmt.Println(a != 5 || b == 100)      // -> false
	fmt.Println(!(a > 0))                // -> false
	fmt.Println(!(a == 1) || (b == 100)) // -> true

	// OPERATORS FOR POINTERS (&) AND CHANNELS (<-)

	// we'll come back to these later
}
