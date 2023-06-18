package main

import "fmt"

func main() {
    // EXERCISE 1

    // x, y, z := 10, 15.5, "Gophers"
    // score := []int{10, 20, 30}
    // fmt.Printf("%d\n", x)
    // fmt.Printf("%f\n", y)
    // fmt.Printf("%s\n", z)
    // fmt.Printf("%v\n", score)

    // fmt.Printf("%q\n", z)

    // fmt.Printf("%v, %v, %v, %v\n", x, y, z, score)

    // fmt.Printf("%T, %T\n", y, score)

    // EXERCISE 2

    // const x float64 = 1.422349587101
    // fmt.Printf("%.4f\n", x)

    // EXERCISE 3

    shape := "circle"
    radius := 3.2
    const pi float64 = 3.14159
    circumference := 2 * pi * radius

    // fmt.Printf("Shape: %q\n") // Need to pass in the verb:
    fmt.Printf("Shape: %q\n", shape)
    // fmt.Printf("Circle's circumference with radius %d is %b\n", radius, circumference) // radius is float, circumference is float
    fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
    _ = shape
}
