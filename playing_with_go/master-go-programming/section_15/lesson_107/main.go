package main

import "fmt"

func main() {
	// EXERCISE 1

	var m1 map[string]string
	fmt.Printf("Type: %T, Value: %#v\n", m1, m1)

	m2 := map[int]string{1: "Taylor Swift", 2: "The Beatles"}
	m2[10] = "Abba"

	fmt.Printf("Value of %d: %q\n", 1, m2[1])
	fmt.Printf("Value of %d: %q\n", 3, m2[3])

	// EXERCISE 2

	// var m3 map[int]bool
	// m3[5] = true // may fail... m3 is nil.. instead:
	m3 := map[int]bool{}
	m3[5] = true

	m4 := map[int]int{3: 10, 4: 40}
	m5 := map[int]int{3: 10, 4: 40}

	// fmt.Println(m4 == m5) // cannot compare maps together like this

	equal := true
	for k, v := range m4 {
		if m5[k] == v {
			continue
		} else {
			equal = false
			break
		}
	}
	fmt.Printf("The two arrays are equal: %v\n", equal)

	// EXERCISE 3

	// m := map[int]bool{"1": true, 2: false, 3: false} // "1" is a string
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 2)

	for k, v := range m {
		fmt.Println(k, "->", v)
	}
}
