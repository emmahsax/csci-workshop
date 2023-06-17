package main

import "fmt"

func main() {
    var age int = 30 // The "int" is optional as type can be inferred
    fmt.Println("Age:", age)

    var name = "Dan"
    fmt.Println("Your name is:", name)
    _ = name // Bypass variable declared and not used, by using _ instead

    s := "Learning golang!"
    fmt.Println(s)

    // Get an error for this because it's not a new declaration:
    //   name := "Andrei"
    // But not for this:
    name = "John"
    _ = name
}
