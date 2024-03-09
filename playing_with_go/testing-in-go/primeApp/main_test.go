package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// Format method starting with "Test" and then any character that isn't a lowercase letter

func Test_checkNumbers(t *testing.T) {
	inputTests := []struct {
		name     string
		input    string
		expected string
	}{
		{"negative", "-1", "-1 is negative, and negative numbers are not prime, by definition!"},
		{"empty", "", "Please enter a whole number!"},
		{"zero", "0", "0 is not prime, by definition!"},
		{"one", "1", "1 is not prime, by definition!"},
		{"two", "4", "4 is not a prime number because it is divisible by 2!"},
		{"prime", "7", "7 is a prime number!"},
		{"typed", "three", "Please enter a whole number!"},
		{"decimal", "1.1", "Please enter a whole number!"},
		{"quit", "q", ""},
		{"QUIT", "Q", ""},
		{"non-latin", "ÿæŵñį", "Please enter a whole number!"},
	}

	for _, e := range inputTests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("Incorrect value returned: expected %q, but got %q", res, e.expected)
		}
	}
}

func Test_intro(t *testing.T) {
	// Save a copy of os.Stdout
	oldOut := os.Stdout

	// Create a read and write pipe
	r, w, _ := os.Pipe()

	// Set os.Stdout to the write pipe
	os.Stdout = w

	intro()

	// Close the writer
	_ = w.Close()

	// Reset os.Stdout to what it was before
	os.Stdout = oldOut

	// Read the output from prompt() function
	out, _ := io.ReadAll(r)

	// Perform test
	if !strings.Contains(string(out), "Enter a whole number, and we'll tell you if it's prime or not. Enter q to quit.") {
		t.Errorf("Incorrect intro text: got %q", string(out))
	}
}

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"negative number", -3, false, "-3 is negative, and negative numbers are not prime, by definition!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true, but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false, but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s, but  got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// Save a copy of os.Stdout
	oldOut := os.Stdout

	// Create a read and write pipe
	r, w, _ := os.Pipe()

	// Set os.Stdout to the write pipe
	os.Stdout = w

	prompt()

	// Close the writer
	_ = w.Close()

	// Reset os.Stdout to what it was before
	os.Stdout = oldOut

	// Read the output from prompt() function
	out, _ := io.ReadAll(r)

	// Perform test
	if string(out) != "\n-> " {
		t.Errorf("Incorrect prompt text: expected \"\\n-> \", but got %q", string(out))
	}
}

func Test_readUserInput(t *testing.T) {
	// Need a channel and an instance of io.Reader
	doneChan := make(chan bool)

	// Create a reference to a bytes.Buffer
	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
