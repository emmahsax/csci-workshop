package main

var taskDone bool // This is ok, since it's a larger scope

func main() {
	// * A name begins with a letter or underscore and may have any number of addtional letters, digits underscores
	// * Case matters! Go has 25 keywords that can only be used where syntax permits
	// 	* https://go.dev/ref/spec#Keywords
	// * Best practice in Go is to use short abbreviations, versus long fully written variables

	// Good
	var mv int // Stands for max value
	var n int

	// Bad
	var maxValue int   // Not ideal, as name is too long, but acceptable
	var max_value int  // Name uses snakecase, not camelcase
	var MAX_VALUE int  // Uppercase is ok, but not underscores
	var writeToDB bool // Acceptable, as it's idiomatic
	var writeToDb bool // Not ok, as not idiomatic

	// Uppercase first letters means it'll be exported, so it's special... only use uppercase when it needs to be used
	// in other files or packages

	_, _, _, _, _, _, _ = mv, n, maxValue, max_value, MAX_VALUE, writeToDB, writeToDb
}
