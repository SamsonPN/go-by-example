package main

import (
	"fmt"
	"math"
)

// interfaces are named collections of method signatures

// basic interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}

// implementing this interface on rect and circle types

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// to implement an interface, we just need to implement all methods in the interface
// so we implemented both area() and perim() for the rect struct
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// also implemented geometry interface for the
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// if a variable has an interface type, we can call the methods
// in the named interface
// this takes in any argument of interface type geometry
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// since both circle and rect implement the geometry interface
	// we can use instances of these structs as arguments for
	// the measure function
	measure(r)
	measure(c)
}
