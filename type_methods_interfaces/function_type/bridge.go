package main

import (
	"fmt"
	"strings"
)

// In Go, function types and interfaces complement each other and can work together seamlessly.
//
//	This concept, where function types act as a bridge to interfaces, enables you to treat functions as first-class
//
// citizens while still adhering to Go's interface-oriented design.

//Function types can be used to satisfy interfaces, allowing functions to serve as implementations for interface methods.
//	This is particularly useful when you want to create lightweight implementations or inject behavior dynamically.

type StringProcessor interface {
	Process(s string) string
}

type StringFunc func(s string) string

func (sf StringFunc) Process(s string) string {
	return sf(s) // Call the function
}

func main() {
	// Use a simple function to process strings
	toUpper := StringFunc(func(s string) string {
		return strings.ToUpper(s)
	})

	// Use a different function
	trimSpaces := StringFunc(func(s string) string {
		return strings.TrimSpace(s)
	})

	var processor StringProcessor

	// Assign and use the `toUpper` implementation
	processor = toUpper
	fmt.Println(processor.Process("hello, world")) // Output: HELLO, WORLD

	// Assign and use the `trimSpaces` implementation
	processor = trimSpaces
	fmt.Println(processor.Process("   hello, world   ")) // Output: hello, world
}

//Function Types as Adapters:
//
//The StringFunc type is essentially a function, but with an attached method (Process) that adapts it to the StringProcessor interface.
//This allows any function with the signature func(string) string to be treated as a StringProcessor.
//
//Dynamic Behavior:
//
//You can swap out the function dynamically, injecting different behaviors without changing the interface or other code.
