package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Print a welcome message
	intro()

	// Create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)

	// Start a goroutine to read user input and run program
	go readUserInput(os.Stdin, doneChan)

	// Block until the doneChan gets a value
	<-doneChan

	// Close the channel
	close(doneChan)

	// Say goodbye
	fmt.Println("\nGoodbye.")
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// Read user input
	scanner.Scan()

	// Check to see if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// Try to convert what user typed into an int
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func intro() {
	fmt.Println("Is It Prime?")
	fmt.Println("------------")
	fmt.Println("\nEnter a whole number, and we'll tell you if it's prime or not. Enter q to quit.")
	prompt()
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// Negative numbers are not prime
	if n < 0 {
		return false, fmt.Sprintf("%d is negative, and negative numbers are not prime, by definition!", n)
	}

	// Use the modulus operator to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d!", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func prompt() {
	fmt.Printf("\n")
	fmt.Print("-> ")
}
