package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := string(99) // Converts to string of rune, not string of digits as expected
    fmt.Println(s)  // -> 99 (ascii character for symbol c)

    // s1 := string(44.2) // Error as it cannot convert

    var myStr = fmt.Sprintf("%f", 44.2)
    fmt.Println(myStr) // -> 44.200000
    var myStr1 = fmt.Sprintf("%d", 34)
    fmt.Println(myStr1) // -> 34

    fmt.Println(string(34234))            // -> 薺 (unicode character at that integer)
    fmt.Println(fmt.Sprintf("%d", 34234)) // -> 34234

    s1 := "3.123"
    fmt.Printf("%T\n", s1) // -> string

    var f1, err = strconv.ParseFloat(s1, 64)
    _ = err
    fmt.Printf("%.4f\n", f1*3.4) // -> 10.6182

    i, err := strconv.Atoi("-50") // ascii to int (string to int)
    s3 := strconv.Itoa(20)        // int to ascii (int to string)
    fmt.Printf("i type is %T, i value is %v\n", i, i)
    fmt.Printf("s3 type is %T, s3 value is %s\n", s3, s3)
}
