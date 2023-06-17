package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255 // Maximum value of uint8
	x++               // This overflows at runtime, so Go catches it and turns it to minimum value of uint8 during runtime
	fmt.Println(x)    // -> 0

	// a := int8(255 + 1) // Evaluated at compile time, so Go catches it right away

	var b int8 = 127
	b++                   // Overflows at runtime, so it becomes minimum value of int8
	fmt.Printf("%d\n", b) // -> -128
	b--                   // This will go to its maximum value when it "underflows"
	fmt.Printf("%d\n", b) // -> 127

	f := float32(math.MaxFloat32)
	fmt.Println(f)
	f = f * 1.2 // Float overflows to infinite
	fmt.Println(f)

	// const i int8 = 300 // Compile time error
}
