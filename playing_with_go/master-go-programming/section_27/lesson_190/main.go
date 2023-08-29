package main

import "fmt"

func power(i int, c chan int) {
	c <- i * i
}

func main() {
	// EXERCISE 1

	var c1 chan float64
	c2 := make(<-chan rune)
	c3 := make(chan<- rune)
	c4 := make(chan int, 10)
	fmt.Printf("Channel %v is of type %T\n", c1, c1)
	fmt.Printf("Channel %v is of type %T\n", c2, c2)
	fmt.Printf("Channel %v is of type %T\n", c3, c3)
	fmt.Printf("Channel %v is of type %T\n", c4, c4)

	// EXERCISE 2

	ch := make(chan string)
	go func(s string, c chan string) {
		c <- s
	}("hello", ch)
	fmt.Println("Received from channel:", <-ch)

	// EXERCISE 3

	c := make(chan int)

	go func(n int) {
		c <- n
	}(100)

	fmt.Println(<-c)

	// EXERCISE 4

	chh := make(chan int)
	for i := 1; i <= 50; i++ {
		go power(i, chh)
		fmt.Printf("The square value of %d is %d\n", i, <-chh)
	}

	// EXERCISE 5

	chhh := make(chan int)
	for i := 1; i <= 50; i++ {
		go func(i int) {
			chhh <- i * i
		}(i)
		fmt.Printf("The square value of %d is %d\n", i, <-chhh)
	}
}
