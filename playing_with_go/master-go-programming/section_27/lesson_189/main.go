package main

import (
	"fmt"
	"math"
	"sync"
)

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	wg.Done()
}

func sum(n1 float64, n2 float64, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%.2f\n", n1+n2)
}

func deposit(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b += n
	m.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b -= n
	m.Unlock()
	wg.Done()
}

func main() {
	// EXERCISE 1

	var wg1 sync.WaitGroup
	wg1.Add(1)
	go sayHello("Mr. Wick", &wg1)
	wg1.Wait()

	// EXERCISE 2

	var wg2 sync.WaitGroup
	wg2.Add(3)
	go sum(92.142, 492.4092, &wg2)
	go sum(2.3142, 9547.394092, &wg2)
	go sum(12.142, 942.42, &wg2)
	wg2.Wait()

	// EXERCISE 3

	var wg3 sync.WaitGroup
	wg3.Add(1)
	go func(n float64) {
		defer wg3.Done()
		fmt.Printf("The square root of %.2f is %.4f\n", n, math.Sqrt(n))
	}(64.00)
	wg3.Wait()

	// EXERCISE 4

	var wg4 sync.WaitGroup
	wg4.Add(50)
	for i := 100; i <= 149; i++ {
		go func(n float64) {
			defer wg4.Done()
			fmt.Printf("The square root of %.2f is %.4f\n", n, math.Sqrt(n))
		}(float64(i))
	}
	wg4.Wait()

	// EXERCISE 5

	var wg sync.WaitGroup
	wg.Add(200)
	balance := 100

	var m sync.Mutex // Adding a mutex

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &m)
		go withdraw(&balance, i, &wg, &m)
	}
	wg.Wait()
	fmt.Println("Final balance value:", balance)
}
