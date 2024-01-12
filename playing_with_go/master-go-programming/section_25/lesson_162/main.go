package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	brand     string
	licenseNo string
}

func (c car) License() string {
	return c.licenseNo
}

func (c car) Name() string {
	return c.brand
}

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	// EXERCISE 1

	var vehicle vehicle = car{brand: "Toyota", licenseNo: "931KEW"}

	fmt.Println(vehicle.License())
	fmt.Println(vehicle.Name())
	fmt.Printf("%T\n", vehicle)

	// EXERCISE 2

	var empty interface{}
	empty = 3
	fmt.Printf("%T, %v\n", empty, empty)

	empty = 292.1498
	fmt.Printf("%T, %v\n", empty, empty)

	empty = []int{4, 29, 409}
	fmt.Printf("%T, %v\n", empty, empty)

	// EXERCISE 3

	var x interface{} = cube{edge: 5}
	fmt.Printf("%T\n", x)
	v := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v)
}
