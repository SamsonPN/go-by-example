package main

import (
	"fmt"
	"math"
	"reflect"
)

// constants can be declared anywhere a var can
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	// constants that are numbers do not have a type initially
	// until they are explicitly converted to one!!!
	// if you hover of "d", you'll see it as untyped float!
	fmt.Println(int64(d))

	fmt.Println("$reflect.TypeOf(float32(d))")

	// if you hover over e, you'll see it as integer type
	e := 123
	fmt.Println(reflect.TypeOf(e))

	// it can be given a type if you use it in a context in which it needs one
	// such as when you can this math.Sin method here on the constant, n
	fmt.Println(math.Sin(n))
}
