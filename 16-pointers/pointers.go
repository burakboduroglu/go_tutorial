package main

import "fmt"

func main() {
	// pointer is a variable that stores the memory address of another variable

	// declare a variable
	var a int = 10

	// declare a pointer variable
	var p *int

	// assign the address of a to p
	p = &a

	fmt.Println("Value of a:", a)
	fmt.Println("Value of *p:", *p)
	fmt.Println("Address of a:", &a)
	fmt.Println("Address of p:", p)

	// change the value of a through the pointer
	*p = 20
}

// Output:
// Value of a: 10
// Address of a: 0xc0000b6010
// Value of p: 0xc0000b6010
// Value of *p: 10
//
// The value of a is 10, the address of a is 0xc0000b6010, the value of p is 0xc0000b6010, and the value of *p is 10.
// The value of a is changed to 20 through the pointer p.
