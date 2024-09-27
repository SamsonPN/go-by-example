package main

import "fmt"

// accepts a variable number of integers
// and puts them into nums
// which would be an array of integers
// so var nums []int
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	// index, value
	for i, num := range nums {
		fmt.Println("i:", i)
		total += num
	}

	fmt.Println("total:", total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// if you have multiple values in a slice
	// you can put them all into a variadic function this way
	// similar to the spread syntax in JS but after the var name
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
