package main

import "fmt"

type duration int

func main() {
	// EXERCISE 1

	// var hour duration
	// fmt.Println(hour)
	// fmt.Printf("%T\n", hour)
	// hour = 3600
	// fmt.Println(hour)

	// EXERCISE 2

	// var hour duration = 3600
	// minute := 60
	// // fmt.Println(hour != minute) // Mismatched types, since hour is type duration and minute is type int
	// fmt.Println(hour != duration(minute))

	// EXERCISE 3

	type mile float64
	type kilometer float64
	const m2km = 1.609

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)

	fmt.Println("The distance in km from Berlin to Paris is", kmBerlinToParis)
}
