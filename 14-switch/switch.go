package main

import "fmt"

func main() {
	// switch case
	// switch case is alternative to if else
	// switch case is used to compare the value of a variable to multiple values

	var number = 2
	switch number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown Number")
	}
}

/*
	Output:
	Two
*/
