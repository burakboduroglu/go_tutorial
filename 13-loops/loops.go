package main

import "fmt"

func main(){
	// Loop-1
	fmt.Println("Loop-1")
	for i:=0; i<5; i++ {
		fmt.Printf("%d\n", i)
	}
	fmt.Println("**********")

	// Loop-2
	fmt.Println("Loop-2")
	i := 0
	for i<5 {
		fmt.Printf("%d\n", i)
		i++
	}
	fmt.Println("**********")

	// Loop-3
	// break means exit the loop
	fmt.Println("Loop-3")
	i = 0
	for {
		if i>4 {
			break
		}
		fmt.Printf("%d\n", i)
		i++
	}
	fmt.Println("**********")

	// Loop-4
	// continue means skip the current iteration
	fmt.Println("Loop-4")
	for i:=0; i<5; i++ {
		if i==2 {
			continue
		}
		fmt.Printf("%d\n", i)
	}
	fmt.Println("**********")

	// Loop-5
	// Nested loop
	fmt.Println("Loop-5")
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			fmt.Printf("%d-%d\n", i, j)
		}
	}
	fmt.Println("**********")
}

/*
	Output:
	Loop-1
	0
	1
	2
	3
	4
	**********
	Loop-2
	0
	1
	2
	3
	4
	**********
	Loop-3
	0
	1
	2
	3
	4
	**********
	Loop-4
	0
	1
	3
	4
	**********
	Loop-5
	0-0
	0-1
	0-2
	1-0
	1-1
	1-2
	2-0
	2-1
	2-2
	**********
*/