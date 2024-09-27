package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// anonymous functions can also be used for recursion
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// since fib was declared in main earlier
		// Go knows which function to call with fib here
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
