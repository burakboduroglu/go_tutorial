package main

import "fmt"

func main() {
	user := "John Doe"
	var scanNumber int

	fmt.Print("Enter a number: ")
	// Scan the user input and store it in the scanNumber variable
	// & is used to get the memory address of the variable
	fmt.Scan(&scanNumber)

	fmt.Println("User:", user)
	fmt.Println("Your Number:", scanNumber)
}

/* 
	Output:
	Enter a number: 10
	User: John Doe
	Your Number: 10
*/