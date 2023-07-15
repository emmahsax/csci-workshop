package main

import "fmt"

type Person struct {
	name           string
	age            int
	favoriteColors []string
}

type Grade struct {
	grade  int
	course string
}

type Person2 struct {
	name           string
	age            int
	favoriteColors []string
	grade          Grade
}

func main() {
	// EXERCISE 1

	you := Person{name: "You", age: 50, favoriteColors: []string{"blue", "green"}}
	me := Person{name: "Me", age: 20, favoriteColors: []string{"yellow", "purple"}}

	fmt.Println(you)
	fmt.Println(me)

	// EXERCISE 2

	me.name = "Andrei"
	newColors := you.favoriteColors

	fmt.Println(newColors)

	you.favoriteColors = append(you.favoriteColors, "black")

	fmt.Println(you)
	fmt.Println(me)

	// EXERCISE 3

	for _, color := range me.favoriteColors {
		fmt.Println(color)
	}

	// EXERCISE 4

	you2 := Person2{
		name:           "You",
		age:            50,
		favoriteColors: []string{"blue", "green"},
		grade: Grade{
			grade:  94,
			course: "Go",
		},
	}

	fmt.Println(you2)
}
