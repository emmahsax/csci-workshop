package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// EXERCISE 1

	var name string = "Emma"
	country := "United States"
	fmt.Printf("Your name: `%s`\nCountry: `%s`\n", name, country)

	fmt.Println("He says: \"Hello\"")
	fmt.Println(`C:\Users\a.txt`)

	// EXERCISE 2

	r := 'ă'
	fmt.Printf("%T\n", r)

	str1 := "ma"
	str2 := "m"
	str3 := str1 + str2 + string(r)
	fmt.Println(str3)

	// EXERCISE 3

	s1 := "țară means country in Romanian"

	fmt.Printf("Bytes in string: ")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v ", s1[i])
	}
	fmt.Println("")

	fmt.Printf("Byte characters in string: ")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c ", s1[i])
	}
	fmt.Println("")

	fmt.Printf("Runes with bytes in string:\n")
	for i, v := range s1 {
		fmt.Printf("%d -> %c\n", i, v)
	}

	// EXERCISE 4

	s2 := "Go is cool!"
	x := s2[0] // -> 71 (byte at G)
	fmt.Println(x)

	// s2[0] = 'x' // s2[0] is a byte, giving it a rune
	// Strings are immutable
	s2 = strings.Replace(s2, "G", "x", 1)

	// printing the number of runes in the string
	fmt.Println(len(s2))                    // -> 11 (returns number of bytes, which happens to be the length as well)
	fmt.Println(utf8.RuneCountInString(s2)) // -> 11 (what we actually want)

	fmt.Println(s2)

	// EXERCISE 5

	s := "你好 Go!"
	ss := []byte(s)

	fmt.Printf("%#v\n", ss)

	for i, v := range ss {
		fmt.Printf("%d -> %d\n", i, v)
	}

	// EXERCISE 6

	rs := []rune(s)

	fmt.Printf("%#v\n", rs)

	for i, v := range rs {
		fmt.Printf("%d -> %c\n", i, v)
	}
}
