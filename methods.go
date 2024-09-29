package main

import "fmt"

// go supports methods defined on struct types

type rect struct {
	width, height int
}

// this function has a receiver type before the func name
// this signifies what which struct uses this method
func (r *rect) area() int {
	return r.width * r.height
}

// methods can be defined with either pointer or value receivers
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r

	// go handles automatic conversions between from pointers to values
	// for method calls
	// e.g. area receives a pointer of rect while perim receives the value

	// why the distinction?
	// you use the pointer when you want to mutate the struct directly
	// or when you do not want to copy the value
	fmt.Println("area:", rp.area())
	fmt.Println("area:", rp.perim())
}
