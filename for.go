package main

import "fmt"

func main() {

	// single condition for-loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\n")

	// basic for-loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("\n")

	// using range
	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Println("\n")

	// infinite loop without a condition
	// unless you break or return from the enclosing function
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("\n")

	// continue keyword
	for i := range 6 {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}

}
