package main

import (
	"fmt"
	"time"
)

func main() {
	// basic switch statement
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// can use commas to separate multiple expressions in the same case statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch without an expression is another way of writing if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// a type switch compares types instead of values
	// can use this to find type of an interface value

	// interface{} is similar to "any" in TypeScript
	// it allows the function to take any value as an argument
	whatAmI := func(i interface{}) {
		// THIS IS A type Switch THAT CAN ONLY BE USED IN SWITCH STATEMENTS!!!
		// you cannot use this anywhere else.

		// type assertions, i.e. i.(type) can be used in if-statements
		// but only if i is of type interface{} like so: var i interface{} = 3
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
