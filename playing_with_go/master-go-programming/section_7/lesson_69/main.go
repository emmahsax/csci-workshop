package main

import "fmt"

func main() {
	// EXERCISE 1

	for n := 1; n <= 50; n++ {
		if n%7 == 0 {
			fmt.Println(n)
		}
	}
	fmt.Println("")

	// EXERCISE 2

	for n := 1; n <= 50; n++ {
		if n%7 != 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Println("")

	// EXERCISE 3

	counter := 0
	for n := 1; n <= 50; n++ {
		if n%7 != 0 {
			continue
		}

		fmt.Println(n)
		counter++

		if counter >= 3 {
			break
		}
	}
	fmt.Println("")

	// EXERCISE 4

	for x := 0; x <= 500; x++ {
		if x%7 == 0 && x%5 == 0 {
			fmt.Println(x)
		}
	}
	fmt.Println("")

	// EXERCISE 5

	birthYear := 2008
	currentYear := 2023

	for y := birthYear; y <= currentYear; {
		fmt.Println(y)
		y++
	}
	fmt.Println("")

	// EXERCISE 6

	// // Original code
	// age := -9
	// if age < 0 || age > 100 {
	//     fmt.Println("Invalid Age")
	// } else if age < 18 {
	//     fmt.Println("You are minor!")
	// } else if age == 18 {
	//     fmt.Println("Congratulations! You've just become major!")
	// } else {
	//     fmt.Println("You are major!")
	// }

	// New code
	age := -9
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
