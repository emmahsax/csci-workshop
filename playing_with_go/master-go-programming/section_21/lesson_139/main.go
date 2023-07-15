package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cube(a float64) float64 {
	return a * a * a
}

func f1(n uint) (uint, uint) {
	var factorial uint
	var sum uint
	factorial = 1
	sum = 0

	for i := n; i > 0; i-- {
		factorial *= i
		sum += i
	}

	return factorial, sum
}

func myFunc(n string) int {
	nn := strings.Repeat(n, 2)
	nnn := strings.Repeat(n, 3)
	i, _ := strconv.Atoi(n)
	ii, _ := strconv.Atoi(nn)
	iii, _ := strconv.Atoi(nnn)
	return i + ii + iii
}

func sum(a ...int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}

func sum2(a ...int) (sum int) {
	sum = 0
	for _, i := range a {
		sum += i
	}
	return
}

func searchItem(a []string, b string) bool {
	val := false

	for i := 0; i < len(a); i++ {
		if a[i] == b {
			val = true
		}
	}

	return val
}

func searchItem2(a []string, b string) bool {
	val := false

	for i := 0; i < len(a); i++ {
		// if strings.ToLower(a[i]) == strings.ToLower(b) {
		// 	val = true
		// }

		// I got a suggestion to use strings.EqualFold instead of strings.ToLower twice:
		if strings.EqualFold(a[i], b) {
			val = true
		}
	}

	return val
}

func print(msg string) {
	fmt.Println(msg)
}

func main2() {
	// print("The Go gopher is the iconic mascot of the Go project.") // Changing this line to add defer first
	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
}

func main() {
	// EXERCISE 1

	fmt.Println(cube(4))

	// EXERCISE 2

	fmt.Println(f1(5))

	// EXERCISE 3

	fmt.Println(myFunc("5"))

	// EXERCISE 4

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	// EXERCISE 5

	fmt.Println(sum2(1, 2, 3))
	fmt.Println(sum2(1, 2, 3, 4, 5, 6))

	// EXERCISE 6

	fmt.Println(searchItem([]string{"wendy", "john", "charlie", "lucas"}, "charlie"))
	fmt.Println(searchItem([]string{"wendy", "john", "charlie", "lucas"}, "billy"))

	// EXERCISE 7

	fmt.Println(searchItem2([]string{"wEndy", "joHn", "ChaRlie", "LucAs"}, "chArlie"))
	fmt.Println(searchItem2([]string{"wEndy", "joHn", "ChaRlie", "LucAs"}, "biLly"))

	// EXERCISE 8

	main2()

	// EXERCISE 9

	varFunc := func(a int) {
		fmt.Println(a)
	}
	fmt.Printf("%T\n", varFunc)
	varFunc(4)
}
