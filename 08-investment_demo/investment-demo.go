package main

import (
	"fmt"
	"math"
)

func main() {
	var inflationRate float64 = 0.035
	var investment float64 = 1000
	var years float64

	fmt.Print("Enter the number of years: ")
	fmt.Scanf("%f", &years)

	futureValue := investment * math.Pow((1 + inflationRate/100), years)
	fmt.Printf("The future value of the investment is $%.2f\n", futureValue)
}
