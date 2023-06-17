package main

import "fmt"

func main() {
    var x string = "my"
    var y string = "name"
    var z string = "is"
    var a string = "Emma"

    fmt.Println(x, y, z, a)                      // It automatically prints spaces between each variable
    fmt.Println(x + " " + y + " " + z + " " + a) // But here, spaces are required to be defined
}
