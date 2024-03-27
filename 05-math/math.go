package main

import (
	"fmt"
	"math"
)

func main() {
	var radius = 10.0

	// Calculate area and circumference of a circle
	// Area = π * r^2
	var area = math.Pi * (radius * radius)
	
	// Calculate circumference of a circle
	// Circumference = 2 * π * r
	var circumference = 2 * math.Pi * radius
	
	fmt.Println("Radius:", radius)
	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)
}

/*
	Output:
		Radius: 10
		Area: 314.1592653589793
		Circumference: 62.83185307179586
*/

