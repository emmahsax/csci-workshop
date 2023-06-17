package main

import "fmt"

func main() {
	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	// I can reuse variables in this statement with :=, as long as
	// at least one of the variables is new:
	car, year := "BMW", 2018
	fmt.Println(car, year)

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	// Use short declaration (:= or var ____ =) when you know the initial value
	// Use var(...) when you don't know initial value, and for readability

	var i, j int
	i, j = 5, 8
	j, i = i, j
	fmt.Println("i:", i, "j:", j)

	sum := 5 + 2.3
	fmt.Println(sum)
}
