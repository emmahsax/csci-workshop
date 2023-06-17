package main

func main() {
    type age int        // int is its underlying type
    type oldAge age     // int is its underlying type
    type veryOldAge age // int is its underlying type
}
