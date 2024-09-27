package main

import (
	"fmt"
	"unicode/utf8"
)

// Go string = read-only slice of bytes
// the standard library and language treat strings as containers of text
// encoded in UTF-8

// in other langues, strings are made of characters
// in Go, runes are the character equivalent

// rune = an integer that represents a Unicode code point
func main() {

	const s = "สวัสดี"

	// since strings are []byte
	// this will give us the length of the raw bytes stored within
	fmt.Println("Len:", len(s))

	// string[i] gives us the raw byte value
	// i.e. the hex values of all the bytes
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// to count how many runes are in a string, we use the utf8 package
	// the run-time of this method depends on the size of the string
	fmt.Println("rune count:", utf8.RuneCountInString(s))

	// range loop handles strings specially
	// and decodes  each rune along with its offset in the string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRunInString")
	// can achieve the same iteration by using DecodeRuneInString
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}

	name := "Samson"
	for _, char := range name {
		fmt.Printf("%#U\n", char)
	}
}

func examineRune(r rune) {
	// single quotes are for rune literals
	// can compare rune values to rune literals directly
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
