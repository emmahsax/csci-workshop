package main

import (
    "fmt"
)

func main() {
    var x = 3
    var y = 3.1

    // x = x * y // Mis-matched types
    x = x * int(y)
    fmt.Println(x)             // -> 9
    fmt.Printf("%T\n", y)      // -> float64
    fmt.Printf("%T\n", int(y)) // -> int

    x = int(float64(x) * y)
    fmt.Printf("%T\n", x) // -> int
    fmt.Println(x)        // -> 27

    y = float64(x) * y
    fmt.Println(y) // -> 83.7

    var a = 5
    fmt.Printf("%T\n", a) // -> int
    var b int64 = 64
    fmt.Printf("%T\n", b) // -> int64
    q := a == int(b)      // this is ok
    // a == b // this is not ok

    _ = q
}
