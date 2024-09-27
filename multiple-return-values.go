package main

import "fmt"

// this returns 2 integers
func vals() (int, int) {
	return 3, 7
}

func main() {
	// multiple assignment
	// from calling vals with multiple return values
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// able to use blank identifier to omit values we do not want
	_, c := vals()
	fmt.Println(c)
}
