package main

import "fmt"

func main() {
	// Example 1
	var integer1 = 10 
	var float1 = 12.5

	// Convert integer to float64
	var result = float64(integer1) + float1 

	// /n means new line
	fmt.Println("Example-1\nResult:", result)

	// Example 2
	// Declare float2 as float64 and assign 10 to it
	var float2 float64 = 10
	var float3 = 12.5

	var result2 = float2 + float3

	fmt.Println("Example-2\nResult 2:", result2)
}

/*
	Output:
		Example-1
		Result: 22.5
		Example-2
		Result 2: 22.5
*/