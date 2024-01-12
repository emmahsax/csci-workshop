package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type object interface {
	volume() float64
}

type geometry interface {
	getColor() string
	object
	shape
}

type cube struct {
	color string
	edge  float64
}

func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)
}

func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

func (c cube) getColor() string {
	return c.color
}

func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v
}

func main() {
	c := cube{edge: 2}
	a, v := measure(c)
	fmt.Printf("Area: %v, Volume: %v\n", a, v)
}

// ./main.go:57:6: invalid recursive type interface1
//         ./main.go:57:6: interface1 refers to
//         ./main.go:62:6: interface2 refers to
//         ./main.go:67:6: interface3 refers to
//         ./main.go:57:6: interface1

// type interface1 interface {
// 	interface2
// 	method1()
// }

// type interface2 interface {
// 	interface3
// 	method2()
// }

// type interface3 interface {
// 	interface1
// 	method3()
// }
