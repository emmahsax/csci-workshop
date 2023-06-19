package main

import "fmt"

func main() {
	// EXERCISE 1

	var cities [2]string
	fmt.Println(cities)

	grades := [3]float64{0: 94.6, 1: 89.2}
	fmt.Println(grades)

	var taskDone = [...]bool{true, false, false}
	fmt.Println(taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Printf("%s", cities[i])
	}

	for index, _ := range grades {
		fmt.Printf("index: %d, value: %.2f\n", index, grades[index])
	}

	// EXERCISE 2

	nums := [...]int{30, -1, -6, 90, -6, 7}
	var count int = 0

	for _, value := range nums {
		if value > 0 && value%2 == 0 {
			count++
		}
	}
	fmt.Printf("We found %d positive and even numbers\n", count)

	// EXERCISE 3

	myArray := [3]float64{1.2, 5.6}
	fmt.Println(myArray)

	myArray[2] = 6
	fmt.Println(myArray)

	a := 10
	// myArray[0] = a // a is of type int, myArray has float64s
	myArray[0] = float64(a)
	fmt.Println(myArray)

	// myArray[3] = 10.10 // Out of bounds
	myArray[2] = 10.10
	fmt.Println(myArray)
}
