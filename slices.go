package main

import (
	"fmt"
	"slices"
)

func main() {
	// if slices are not uninitialized, they are assigned nil
	// b/c they are typed ONLY by the elements they contain
	// and are of length 0
	// very different from arrays where it is initialized with zero-values
	// according to the number of elements they contain
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// this allows us to make an empty slice of length 3
	// similar to arrays, using make does initialize it with zero-values
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// can set/get like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// append returns a new slice
	// with the new addedvalues
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// creating an empty slice c of where len(c) == len(s)
	// and then copying s into c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slices also support a slice operator
	// slice[low:high] exactly like python
	// [low, high)

	l := s[2:5]
	fmt.Println("sl1", l)

	// [0, 5)
	l = s[:5]
	fmt.Println("sl2:", l)

	// [2, len(s))
	l = s[2:]
	fmt.Println("sl3:", l)

	// can declare and initialize a slice similar to an array
	// but without specifying the size
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// the slices package has a lot of useful methods inside it
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// slices can also be multi-dimensional
	// the length of the inner slices can vary unlike with
	// multi-dimensional arrays

	// outer slice is of length 3
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
