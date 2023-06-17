package main

import "fmt"

type km float64
type mile float64

func main() {
    type speed uint

    var s1 speed = 10
    var s2 speed = 20

    fmt.Println(s2 - s1) // -> 10

    var x uint
    // x = s1 // Error, different types
    x = uint(s1)
    fmt.Println(x) // -> 10

    var s3 speed = speed(x)
    fmt.Println(s3) // -> 10

    var parisToLondon km = 470
    var distanceInMiles mile

    // distanceInMiles = parisToLondon / 0.621 // Error, different types, since one is km and one is mile
    distanceInMiles = mile(parisToLondon) * 0.621
    fmt.Println(distanceInMiles)
}
