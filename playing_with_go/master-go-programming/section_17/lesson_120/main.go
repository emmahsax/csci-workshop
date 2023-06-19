package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// EXERCISE 1

	file, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	// EXERCISE 2

	_, err = os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}

	os.Rename("info.txt", "newInfo.txt")

	// EXERCISE 3

	err = os.Remove("newInfo.txt")
	if err != nil {
		log.Fatal(err)
	}

	// EXERCISE 4

	// os.ReadFile will open and close the file for you
	// It's a more updated alternative to ioutil.ReadAll()
	readFile, err := os.ReadFile("test.md")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("readFile type: %T\n", readFile)
	fmt.Println("File as bytes:", readFile)
	fmt.Println("File as string:", string(readFile))

	// EXERCISE 5

	fmt.Println(strings.Repeat("#", 100))
	fmt.Println("")

	file, err = os.Open("test.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// EXERCISE 6

	file, err = os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	_, err = bufferedWriter.WriteString("The Go gopher is an iconic mascot!")
	if err != nil {
		log.Fatal(err)
	}
	bufferedWriter.Flush()
}
