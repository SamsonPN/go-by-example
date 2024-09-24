package main

import "fmt"

func main() {
	// THERE IS NO TERNARY OPERATOR IN GO!!!

	// just typical if-statements
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// statements can precede conditionals
	// as long as you have a semi-colon that separates them
	// and it is available in the rest of the if-block

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
