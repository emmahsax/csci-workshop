package main

import (
	"fmt"
)

func main() {
	// EXERCISE 1

	// s := []string{"Go", "for", "Go!"}

	// for i, v := range s {
	//     fmt.Printf("Index: %d, Value: %s\n", i, v)
	// }

	// EXERCISE 2

	// mySlice := []float64{1.2, 5.6}
	// fmt.Println(mySlice)

	// // mySlice[2] = 6 // Index out of range runtime error
	// mySlice[1] = 6.0
	// fmt.Println(mySlice)

	// a := 10
	// // mySlice[0] = a // Mismatched type error
	// mySlice[0] = float64(a)
	// fmt.Println(mySlice)

	// mySlice[1] = 10.10 // Index out of range runtime error
	// fmt.Println(mySlice)

	// mySlice = append(mySlice, float64(a))
	// fmt.Println(mySlice)

	// EXERCISE 3

	// nums := []float64{1.2, 3.4, 5.6}
	// nums = append(nums, 10.1)
	// nums = append(nums, 4.1, 5.5, 6.6)
	// fmt.Println("nums:", nums)

	// n := []float64{1.2, 3.4}
	// nums = append(nums, n...)
	// fmt.Println("nums:", nums)

	// EXERCISE 4

	// if len(os.Args) < 3 || len(os.Args) > 11 {
	//     fmt.Println("Please pass in between 2 and 10 numbers")
	//     os.Exit(1)
	// }

	// args := os.Args[1:]
	// fmt.Println(args)
	// var sum int = 0
	// var product int = 1

	// for _, v := range args {
	//     sum1, err1 := strconv.Atoi(v)
	//     if err1 != nil {
	//         fmt.Println("error received", err1)
	//         break
	//     }
	//     product1, err2 := strconv.Atoi(v)
	//     if err2 != nil {
	//         fmt.Println("error received", err2)
	//         break
	//     }

	//     sum = sum + sum1
	//     product = product * product1
	// }

	// fmt.Printf("Sum is %d, Product is %d\n", sum, product)

	// EXERCISE 5

	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum := 0

	for _, v := range nums[2 : len(nums)-2] {
		fmt.Printf("%d,", v)
		sum = sum + v
	}
	fmt.Printf("\nTotal sum: %d\n", sum)

	// EXERCISE 6

	// friends := []string{"Mary", "John", "Paul", "Diana"}
	// myFriends :=  make([]string, len(friends))
	// copy(myFriends, friends)
	// myFriends[0] = "Sally"
	// fmt.Printf("friends: %#v\n", friends)
	// fmt.Printf("myFriends: %#v\n", myFriends)

	// EXERCISE 7

	friends := []string{"Mary", "John", "Paul", "Diana"}
	myFriends := []string{}
	myFriends = append(myFriends, friends...)
	myFriends[0] = "Sally"
	fmt.Printf("friends: %#v\n", friends)
	fmt.Printf("myFriends: %#v\n", myFriends)

	// EXERCISE 8

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := append(years[:3], years[len(years)-3:]...)
	fmt.Printf("%#v\n", newYears)
}
