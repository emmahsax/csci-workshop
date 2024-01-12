/////////////////////////////////
// Intro to Channels
// Go Playground: https://play.golang.org/p/Uc7iiqVeZLL
/////////////////////////////////

package main

import "fmt"

func main() {
	// Declaring a channel of type `chan int`
	var c1 chan int
	fmt.Println(c1) // => nil (its zero value is nil)

	// Initializing the channel
	c1 = make(chan int)
	fmt.Println(c1) // => 0xc000078060 (the channel stores an address)

	// Declaring and initializing a channel at the same time
	c2 := make(chan int)
	_ = c2

	// Declaring and initilizing a RECEIVE-ONLY channel
	c3 := make(<-chan string)

	// Declaring and initilizing a SEND-ONLY channel
	c4 := make(chan<- string)

	fmt.Printf("%T, %T, %T\n", c1, c3, c4) // => chan int, <-chan string, chan<- string

	//** The "arrow" indicates the direction of data flow!! **//

	// Sending a value into the channel
	go func() {
		c1 <- 10
	}()
	// Receiving a value from the channel
	fmt.Println("Value received:", <-c1)

	// Closing a channel
	close(c1)
}
