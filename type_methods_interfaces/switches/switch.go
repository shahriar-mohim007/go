package main

import "fmt"

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("I am an int, value: %d\n", v)
	case string:
		fmt.Printf("I am a string, value: %s\n", v)
	case bool:
		fmt.Printf("I am a bool, value: %v\n", v)
	default:
		fmt.Printf("Unknown type, value: %v\n", v)
	}
}

func main() {
	describe(42)      // Output: I am an int, value: 42
	describe("hello") // Output: I am a string, value: hello
	describe(true)    // Output: I am a bool, value: true
	describe(3.14)    // Output: Unknown type, value: 3.14
}
