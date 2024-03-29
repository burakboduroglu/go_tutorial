/*
	Exercise

	1. Ask for revenue, expenses, and taxes
	2. Calculate earnings before taxes (ebt) and earnings after taxes (profit)
	3. Calculate ratio
	4. Print the results
*/

/*
	Solution
*/

package main

import (
	"fmt"
)

func main(){
	var revenue float64
	var expenses float64
	var taxRate float64

	//Get Revenue
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	// Get Expenses
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	// Get Taxes
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	fmt.Println("*************")
	fmt.Println("EBT: ", ebt)
	fmt.Println("Ratio: ", ratio)
	fmt.Println("Profit: ", profit)
}

/*
	Output:
		Revenue: 1000
		Expenses: 500
		Tax Rate: 20
		*************
		EBT:  500
		Ratio:  1.25
		Profit:  400
*/