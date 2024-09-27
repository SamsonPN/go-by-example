package main

import "fmt"

// pointers allow you to pass references to values and records within the program

// integer parameter
// arguments will be passed to this func BY VALUE
// this basically just gets a copy of ival, not the actual thing
func zeroval(ival int) {
	ival = 0
}

// integer POINTER parameter
// basically passes in the memory address of the value of the argument
func zeroptr(iptr *int) {
	// this dereferences that value
	// and accesses the actual value at that memory address
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// zeroval receives a COPY of i
	// not the actual value of i
	// so if you change it in zeroval, i does not get changed in main
	zeroval(i)
	fmt.Println("zeroval:", i)

	// a pointer points to the memory address of the variable i
	// &[var-name] allows you to do that

	// zeroptr receives the memory address of i
	// which allows it to change i here in main
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
