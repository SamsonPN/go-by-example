package main

import "fmt"

// structs are a typed collection of fields
// useful for grouping data to form records

type person struct {
	name string
	age  int
}

// this creates a new person struct with the given name
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	// go is a garbage collected language
	// you can safely return a pointer to a local variable
	// it will only be cleaned up by the garbage collection when there are no
	// active references to it
	return &p
}

func main() {

	// creates a new struct
	fmt.Println(person{"Bob", 20})

	// can name the fields when initializng struct
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// adding & will get the pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// idiomatic to encapsulate new struct creation in
	// constructor functions
	fmt.Println(newPerson("Jon"))

	// access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// can also use dots with struct pointers
	// the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// if a struct type is only used for a single value
	// we can just use an anonymous struct
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
