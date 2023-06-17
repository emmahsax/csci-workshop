package main

import "fmt"

func main() {
    const (
        c1 = iota
        c2 = iota
        c3 = iota
    )

    fmt.Println(c1, c2, c3)

    // Iota increments automatically, but restarts back at 0
    const (
        c11 = iota
        c22
        c33
    )

    fmt.Println(c11, c22, c33)

    const (
        North = iota
        East
        South
        West
    )

    fmt.Println(West)

    const (
        a = iota*2 + 1 // 0 * 2 + 1 = 1
        b              // 1 * 2 + 1 = 3
        c              // 2 * 2 + 1 = 5
    )

    fmt.Println(a, b, c)

    // x = -2, y = -4, z = -5
    const (
        x = -(iota + 2) // -(0 + 2)
        _               // -(1 + 2)
        y               // -(2 + 2)
        z               // -(3 + 2)
    )

    fmt.Println(x, y, z)
}
