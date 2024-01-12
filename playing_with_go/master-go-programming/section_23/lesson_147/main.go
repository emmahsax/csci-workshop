package main

import "fmt"

func main() {
	// EXERCISE 1

	x1 := 10.10
	fmt.Printf("x's address is %p\n", &x1)

	ptr := &x1
	fmt.Printf("ptr is of type %T and value %v\n", ptr, ptr)
	fmt.Printf("ptr's address is %p and ptr's value's value is %v\n", &ptr, *ptr)

	// EXERCISE 2

	x2, y1 := 10, 2
	ptrx, ptry := &x2, &y1

	z := *ptrx / *ptry

	fmt.Printf("The value of z: %v\n", z)

	// EXERCISE 3

	x3, y2 := 5.5, 8.8
	fmt.Printf("Values before swap - x: %v, y: %v\n", x3, y2)
	swap(&x3, &y2)
	fmt.Printf("Values after swap - x: %v, y: %v\n", x3, y2)
}

func swap(x *float64, y *float64) {
	z := *x
	*x = *y
	*y = z
}
