package main

import "fmt"

func main() {
	// EXERCISE 1

	// var i int = 3
	// var f float64 = 3.2

	// fmt.Printf("i is now type %T, value is %f\n", float64(i), float64(i))
	// fmt.Printf("f is now type %T, value is %d\n", int(f), int(f))

	// EXERCISE 2

	// var i = 3
	// var f = 3.2
	// var s1, s2 = "3.14", "5"

	// i1 := fmt.Sprint(i) // Could also do strconv.Itoa(i)
	// s21, err1 := strconv.Atoi(s2)
	// f1 := fmt.Sprintf("%f", f)
	// s12, err2 := strconv.ParseFloat(s1, 64)

	// _, _ = err1, err2

	// fmt.Printf("i is now type %T, value is %s\n", i1, i1)
	// fmt.Printf("s2 is now type %T, value is %d\n", s21, s21)
	// fmt.Printf("f is now type %T, value is %s\n", f1, f1)
	// fmt.Printf("s1 is now type %T, value is %f\n", s12, s12)

	// EXERCISE 3

	// x, y := 4, 5.1

	// // z := x * y // Type mismatch
	// z := float64(x) * y
	// fmt.Println(z)

	// // a := 5 // Needs to be a float, so add the dot to the end
	// a := 5.
	// b := 6.2 * a
	// fmt.Println(b)

	// EXERCISE 4

	// const sunToEarth int = 149_600_000 * 1000 // Kilometers to earth in meters
	// const speedOfLight int = 299_792_458

	// time := sunToEarth / speedOfLight
	// fmt.Printf("Sun reaches the earth in %d seconds\n", time)

	// EXERCISE 5

	x, y := 0.1, 5
	var z float64

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.

	result1 := x < z && int(x) != int(z)
	fmt.Println(result1)

	result2 := y == 1*5 || int(z) == 0
	fmt.Println(result2)
}
