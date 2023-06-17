package main

import "fmt"

func main() {
    name := "Emma"
    fmt.Println("Hello World")
    fmt.Println("My name is", name)

    a, b := 4, 6
    fmt.Println("Sum:", a+b, ", Mean Value:", (a+b)/2) // Adds whitespace after each argument, adds newline at end

    // Newlines aren't added to Printf, so we need to specify them

    // %d -> int
    fmt.Printf("Your age is %d\n", 21)
    // %f -> float
    fmt.Printf("x is %d, y is %f\n", 5, 6.8)
    fmt.Printf("He says: \"Hello Go!\"\n")

    figure := "Circle"
    radius := 5
    pi := 3.14159

    fmt.Printf("Radius: %d\n", radius)
    fmt.Printf("Radius: %+d\n", radius)
    fmt.Printf("Pi: %f\n", pi)
    // %s -> string
    fmt.Printf("The diameter of %s with radius of %d is %f\n", figure, radius, float64(radius)*2*pi)
    // %q -> adds double quotes to string
    fmt.Printf("This is a %q\n", figure)
    // %v -> any type
    fmt.Printf("The diameter of %v with radius of %v is %v\n", figure, radius, float64(radius)*2*pi)
    // %T -> prints type
    fmt.Printf("figure is of type %T\n", figure)
    fmt.Printf("radius is of type %T\n", radius)
    fmt.Printf("pi is of type %T\n", pi)

    // %t -> bool
    closed := true
    fmt.Printf("file closed: %t\n", closed)

    // %b -> converts int to base 2
    fmt.Printf("%b\n", 55)
    fmt.Printf("%08b\n", 55) // Tells to use 8 bits with leading 0s

    // %x -> convert decimal to hexadecimal (base 16)
    fmt.Printf("%x\n", 100)

    x := 3.4
    y := 6.9

    fmt.Printf("x * y = %f\n", x*y)
    fmt.Printf("x * y = %.3f\n", x*y) // Forces 3 decimal points
}
