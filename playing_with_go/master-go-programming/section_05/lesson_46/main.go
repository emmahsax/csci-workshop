package main

import "fmt"

func main() {
	// EXERCISE 1

	// const daysWeek int = 7
	// const lightSpeed float64 = 299792458
	// const pi float64 = 3.14159

	// EXERCISE 2

	// const (
	//     daysWeek = 7
	//     lightSpeed = 299792458
	//     pi = 3.14159
	// )

	// EXERCISE 3

	// const secPerDay int = 60 * 60 * 24 // seconds/minute, minute/hour, hour/day
	// const daysYear int = 365
	// fmt.Println(daysYear * secPerDay)

	// EXERCISE 4

	// const x int = 10

	// // const m = []int{1: 3, 4: 5, 6: 8} // Slices cannot be constants, must therefore do:
	// var m = []int{1: 3, 4: 5, 6: 8}
	// _ = m

	// EXERCISE 5

	// const a int = 7
	// const b float64 = 5.6
	// // const c = a * b // Mismatched types, cannot combine like this, must therefore do:
	// const c = float64(a) * b
	// // OR can make a an untyped constant: const a = 7 (and then do normal multiplication)

	// x := 8
	// // const xc int = x // x is a variable (runtime), const is compile time, must therefore do:
	// var xc int = x
	// _ = xc

	// // const noIPv6 = math.Pow(2, 128) // math.Pow is a runtime, not compile time, must do:
	// var noIPv6 = math.Pow(2, 128)
	// _ = noIPv6

	// EXERCISE 6

	const (
		_ = iota + 1
		_
		_
		_
		_
		jun
		jul
		aug
		_
		_
		_
		_
	)

	fmt.Println(jun, jul, aug)
}
