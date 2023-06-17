package main

import "fmt"

func main() {
	// ARITHMETIC OPERATORS

	a, b := 4, 2

	r := (a + b) / (a - b) * 2 // Don't forget PERMDAS
	fmt.Println(r)             // -> 6

	r = 9 % a
	fmt.Println(r) // -> 1

	// ASSIGNMENT OPERATORS

	a, b = 2, 3
	a += b
	fmt.Println(a) // -> 5

	b -= 2
	fmt.Println(b) // -> 1

	b *= 10
	fmt.Println(b) // -> 10

	b /= 5
	fmt.Println(b) // -> 2

	a %= 3
	fmt.Println(a) // -> 2

	x := 1
	x += 1
	x++
	fmt.Println(x) // -> 3

	x--
	fmt.Println(x) // -> 2
	// fmt.Println(5 + x--) // Errors out
}
