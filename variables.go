package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// if the variable is declared but not initialized
	// then it is "zero-valued", so for an int, the zero-value is 0
	var e int
	fmt.Println(e)

	// this is an empty string
	var g string
	fmt.Println(g == "")

	/**
	* all the zero-values for common types:
	int: 0
	float64: 0.0
	bool: false
	string: "" (empty string)
	pointer, slice, map, channel, function, interface: nil
	*/

	// this is the shorthand syntax that is only available
	// inside functions
	f := "apple"
	fmt.Println(f)
}
