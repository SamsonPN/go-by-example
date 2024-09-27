package main

import (
	"fmt"
	"maps"
)

func main() {
	// map[key-type]val-type
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// if the key doesn't exist,
	// the zero-value of the type is returned
	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	// len of a map
	// returns number of key/value pairs
	fmt.Println("len: ", len(m))

	// able to delete a key-value pair
	delete(m, "k2")
	fmt.Println("map: ", m)

	// removes all key-value pairs
	clear(m)
	fmt.Println("map: ", m)

	// there's a 2nd return value
	// that tells us if the key IS present
	// basically is the equivalent to map.has()
	// it disambiguates between a key that actually exists in the map but returns 0
	// and a key that does not exist and its value is an integer that returns the zero-value: 0

	// also "_" (underscore) is a blank identifier
	// it's used to denote that we do not need the first return value (the actual value)
	// returned by m["k2"]
	_, prs := m["k2"] // false
	fmt.Println("prs: ", prs)

	// can declare and initialize a new map
	// in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

	// maps package has a bunch of functions for maps
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
