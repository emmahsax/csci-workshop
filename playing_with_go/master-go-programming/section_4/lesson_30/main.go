package main

import "fmt"

func main() {
    // INT TYPE

    var i1 int8 = 100
    fmt.Printf("%T\n", i1)

    // Overflow error
    // var i2 int8 = -129
    // fmt.Printf("%T\n", i2)

    var i2 uint16 = 65535 // This is the maximum of uint16
    fmt.Printf("%T\n", i2)

    // FLOAT TYPE

    var f1, f2, f3 float64 = 1.1, -0.2, .4
    fmt.Printf("%T %T %T\n", f1, f2, f3) // Notice how the 0 before . is added to both f2 and f3

    // RUNE TYPE

    var r rune = 'f'
    fmt.Printf("%T\n", r)
    fmt.Println(r)        // The decimal ascii code of f
    fmt.Printf("%x\n", r) // The hexadecimal ascii code of f

    // BOOL TYPE

    var b bool = true
    fmt.Printf("%T\n", b)

    // STRING TYPE

    var s string = "Hello Go!" // We could omit the types from these since they can be inferred
    fmt.Printf("%T\n", s)
}
