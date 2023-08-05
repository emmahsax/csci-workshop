package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN!\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, _ := io.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1] // e.g. https://www.google.com
			file += ".txt"                      // e.g. www.google.com.txt

			fmt.Printf("Writing response body to %s\n", file)
			err = os.WriteFile(file, bodyBytes, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	wg.Done()
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://www.google.com",
		"https://emmasax.com",
		"https://emmasax.com/this/does/not/exist",
		"https://www.medium.com",
		"https://my.nonexistent.domain.com",
	}

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 60))
	}
	fmt.Println("# of Goroutines:", runtime.NumGoroutine())
	wg.Wait()
}
