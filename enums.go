package main

import "fmt"

// enums = enumerated types
// it's a type that has a fixed number of possible values,each with a unique name
// Go doesn't have native enum types but it can be implemented like so

// our enum type has an underlying int type
type ServerState int

// these are the possible values for ServerState

const (
	// iota = special keyword that generates successive constant values automatically
	// so this would generate 0, 1, 2, etc
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// since ServerState is an enum with all its names being integer types
// you can use their indices to access this string array
// to get their corresponding strings
// var stateName = [...]string{
// 	"idle",
// 	"connected",
// 	"error",
// 	"retrying",
// }

// this is the interface
// when you implement it,
// every time you fmt.Println etc, you call the function below
// type Stringer interface {
//     String() string
// }

// by implementing fmt.Stringer interface,
// the values of ServerState can be printed out or converted to strings
func (ss ServerState) String() string {
	return stateName[ss]
}

// if we have a value of type int, we cannot pass it to transition
// compiler will complain about type mismatch
// this gives us some degree of compile-time type safety for enums
func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
