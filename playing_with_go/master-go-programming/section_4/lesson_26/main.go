package main

func main() {
	// Constant Rules

	// 1. You cannot change a constant

	const temp int = 100
	// temp = 50 // Will error

	// 2. You cannot initialize a constant at runtime... you must know it at compile time

	// const power = math.Pow(2, 3) // Will error

	// 3. You cannot use a variable to initialize a constant, as variables belong to runtime

	t := 5
	// const tc = t // Will error
	_ = t

	// 4. You can use a function to initialize a constant IF the function has a constant or literal constant as an argument
	const l1 = len("Hello") // This is ok
}
