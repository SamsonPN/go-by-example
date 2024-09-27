package main

import "fmt"

// range lets you iterate over elements for a variety of built-in data structures
func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	// range provides both the index and value for each entry for arrays and slices
	// use the blank identifier "_" if you don't need one or the other
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	// range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// can also just iterate over keys of a map
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// range on strings iterates over unicode code points
	// first value is the starting byte index of the rune
	// and second value is the rune itself
	// i think runes are Go's version of characters but i'm not there yet
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
