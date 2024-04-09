package main

import "fmt"

func main(){
	// Call the function sumAndReturn
	result := sumAndReturn(10, 20)
	fmt.Println(result)


	// Call the function sumAndPrint
	sumAndPrint(10, 20)

	// Call the function returnTwoValues
	value1, value2 := returnTwoValues()
	fmt.Println(value1, value2)

	// Call the function returnTwoValuesAlternative
	value1, value2 = returnTwoValuesAlternative()
	fmt.Println(value1, value2)
}


// Function to sum two numbers amd return the result
func sumAndReturn(a int, b int) int {
	result := a + b
	return result
}

// Function to sum two numbers and print the result
func sumAndPrint(a int, b int) {
	result := a + b
	fmt.Println(result)
}

// Return two values from a function
func returnTwoValues() (int, int) {
	return 10, 20
}

// Alternative way to return two values from a function
func returnTwoValuesAlternative() (value1 int, value2 int) {
	value1 = 10
	value2 = 20
	return
}

/*
	Output:
	30
	30
	10 20
	10 20
*/