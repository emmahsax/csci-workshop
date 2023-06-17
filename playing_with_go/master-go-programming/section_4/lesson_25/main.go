package main

import "fmt"

func main() {
    // Constants represent fixed and unchanging values
    // As they belong to compile time, errors are discovered during compilation
    // Basic literals (1, 3.4, "hello", true) are all unnamed constants
    // There are ONLY boolean constants, rune constants, integer constants,
    //   floating-point constants, complex constants, and string constants

    const days int = 7

    var i int // 0 value by default
    fmt.Println(i)

    // This errors out
    // const pi float64
    const pi float64 = 3.14

    const secondsInHour = 3600
    duration := 234 // Hours
    fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

    duration = 12 // Can modify this value
    fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

    x, y := 5, 0 // Variables belong to runtime, so errors are discovered at runtime
    // fmt.Println(x / y) // Errors at runtime, not compilation

    const a = 5
    const b = 0
    // fmt.Println(a / b) // VS Code detects the error when the file is saved

    // Without the below line, the run will stop here... it won't error, but because we
    // haven't used x and y, it soft-stops
    _, _ = x, y
    // Notice how we don't need to blank out constants, only variables!

    const n, m int = 4, 5
    const n1, m1 = 6, 7

    const (
        min1 = -500
        min2 = -300
        min3 = 100
    )

    fmt.Println(min1, min2, min3)

    // All of these instantiate to -500, as min5 and min6 will get their type
    // and value from min4
    const (
        min4 = -500
        min5
        min6
    )

    fmt.Println(min4, min5, min6)
}
