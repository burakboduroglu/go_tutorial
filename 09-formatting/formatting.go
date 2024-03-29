package main

import "fmt"

func main(){
	var floatNumber = 1.4321235232

	// Format the number to 2 decimal places
	fmt.Printf("Number: %.2f\n", floatNumber)
	
	// Format the number to 2 decimal places using Sprintf
	// Sprintf returns a formatted string without printing it
	formatedString := fmt.Sprintf("Number: %.2f\n", floatNumber)
	fmt.Println(formatedString)
}