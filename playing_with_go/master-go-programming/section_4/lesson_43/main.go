package main

import "fmt"

func main() {
	var a uint8 = 10
	var b byte

	b = a          // byte is an alias for uint8
	fmt.Println(b) // -> 10

	type second = uint

	var hour second = 3600
	fmt.Printf("Minutes in an hour: %d\n", hour/60) // -> 60
}
