package main

import "fmt"

func main() {
    const a float64 = 5.1 // This is a typed constant
    const b = 6.7         // This is an untyped constant
    const c float64 = a * b
    const str = "Hello " + "Go!"
    const d = 5 > 10

    fmt.Println(d)

    // This errors out as x and y are different types
    // const x int = 5
    // const y float64 = 2.2 * x

    // This doesn't error out, since they're untyped, BUT it doesn't respond with
    // a correct answer
    const x = 5
    const y = 2.2 * 5
    fmt.Printf("x is %T, y is %T\n", x, y)

    var i int = x     // var i int = x
    var j float64 = x // var j float64 = float64(x)
    var p byte = x    // var p byte = byte(x)

    fmt.Println(i, j, p)

    const r = 5
    var rr = r
    fmt.Printf("%T\n", rr)

    // Untyped constants live in a world that is less restrictive than Go's normal system.
    // But to use them, we need to assign them to variables, and then the variables
    // need to be typed (not the constants themselves).
}
