package main

import (
	"fmt"
)

type money float64

type book struct {
	price float64
	title string
}

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	b.price = b.price * .9
}

func (b *book) changePrice(p float64) {
	b.price = p
}

func main() {
	// EXERCISE 1

	var m money = 1.2849
	fmt.Println(m)
	m.print()

	// EXERCISE 2

	fmt.Println(m.printStr())

	// EXERCISE 3

	b := book{price: 21.48, title: "My Cool Book"}
	fmt.Println("VAT:", b.vat())
	b.discount()
	fmt.Println("Book:", b)

	// EXERCISE 4

	// book value
	bestBook := book{title: "The Trial by Kafka", price: 9.9}

	// changing the price
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price: %.2f\n", bestBook.price) // no change
}
