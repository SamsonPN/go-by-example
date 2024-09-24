package main

import "fmt"

func main() {
	// HAVE TO SPECIFY A LENGTH
	// arrays are situational, use slices (dynamic) instead

	// by default, arrays are zero-valued
	// so for integers, that would mean 0s
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	// java syntax to declare and initialize arrays
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// use ellipses so that the compiler can count len for you
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// can specify the index with ':' and elements between them
	// will be zeroed
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// can have multi-dimensional arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// can create and initialize multi-dimensional arrays too
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("\n")
	fmt.Println("2d: ", twoD)
}
