/////////////////////////////////
// Strings, Runes, Bytes and Unicode Code Points
// Go Playground: https://play.golang.org/p/pttCqLAAvKA
/////////////////////////////////

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	// characters or rune literals are expressed in Go by enclosing them in single quotes
	// declaring a variable of type rune (alias to int32)
	var1 := 'a'
	fmt.Printf("Type: %T, Value:%d \n", var1, var1) // => Type: int32, Value:97
	// value is 97 (the code point for a)

	// declaring a string value that contains non-ascii characters
	str := "ţară" // ţară means country in Romanian
	// 't', 'a' ,'r' and 'a' are runes and each rune occupies beetween 1 and 4 bytes.

	fmt.Printf("String for reference: %s\n", str)

	//The len() built-in function returns the no. of bytes not runes or chars.
	fmt.Println("No. of bytes in string:", len(str)) // -> No. of bytes in string: 6 (4 runes in the string but the byte length is 6)

	// returning the number of runes in the string (what we'd typically think of as string length)
	m := utf8.RuneCountInString(str)
	fmt.Println("No. of runes in string:", m) // => No. of runes in string: 4

	// by using indexes we get the byte at that position, not rune.
	fmt.Println("Byte (not rune) at position 1:", str[1]) // -> 163
	r, _ := utf8.DecodeRuneInString(str[2:])
	fmt.Printf("Rune at position 1: %c\n", r) // -> Rune at position 1: a (need to use index 2 because ţ takes up two bytes)

	// decoding a string byte by byte (the bytes don't match up with the runes)
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d -> %c, ", i, str[i]) // -> Å£arÄ
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	// decoding a string rune by rune manually:
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:]) // it returns the rune in string in variable r
		//and its size in bytes in variable size

		// printing out each rune
		fmt.Printf("%c", r) // -> ţară
		i += size           // incrementing i by the size of the rune in bytes
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	// decoding a string rune by rune automatically:
	for i, r := range str { //the first value returned by range is the index of the byte in string where rune starts
		fmt.Printf("%d -> %c, ", i, r) // => ţară
	}

	my_byte := byte('b')
	my_rune := 'b'

	fmt.Printf("\nByte: %c = %d\nRune: %c = %U\n", my_byte, my_byte, my_rune, my_rune)
	// -> Byte: b = 98
	//    Rune: b = U+0062
}
