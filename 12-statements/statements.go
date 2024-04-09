package main

import "fmt"

func main(){
	checkNumber(5)
	between(15)
}


// Function to check if a number is greater than 10
func checkNumber(number int) {
	if number > 10 {
		fmt.Println("Number is greater than 10")
	} else {
		fmt.Println("Number is less than or equal to 10")
	}
}

// Function to check if a number is between 10 and 20
func between(number int) {
	if (number > 10 && number < 20) {
		fmt.Println("Number is between 10 and 20")
	} else {
		fmt.Println("Number is not between 10 and 20")
	}
}

/*
	Output:
	Number is greater than 10
	Number is not between 10 and 20
*/