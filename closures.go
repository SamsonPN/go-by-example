package main

import "fmt"

// intSeq returns an anonymous function that has access to the variable i
// to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// intSeq() returns a func
	nextInt := intSeq()

	// which you can call here using the var name it was assigned to
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// the state is unique for each function assignment
	// so newInts returns 1 after calling it
	newInts := intSeq()
	fmt.Println(newInts())
}
