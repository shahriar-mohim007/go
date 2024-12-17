package main

import (
	"fmt"
	"time"
)

type LargeStruct struct {
	Data [1024 * 1024]byte // 1MB of data
}

// Pass by value
func processByValue(ls LargeStruct) {
	_ = ls
}

// Pass by pointer
func processByPointer(ls *LargeStruct) {
	_ = ls
}

// Return by value
func createByValue() LargeStruct {
	return LargeStruct{}
}

// Return by pointer
func createByPointer() *LargeStruct {
	return &LargeStruct{}
}

func measureTime(label string, fn func(), iterations int) {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		fn()
	}
	duration := time.Since(start)
	fmt.Printf("%s: %v\n", label, duration)
}

func main() {
	ls := LargeStruct{}

	const iterations = 1000

	// Benchmark passing data
	measureTime("Pass by value", func() { processByValue(ls) }, iterations)
	measureTime("Pass by pointer", func() { processByPointer(&ls) }, iterations)

	// Benchmark returning data
	measureTime("Return by value", func() { _ = createByValue() }, iterations)
	measureTime("Return by pointer", func() { _ = createByPointer() }, iterations)
}

//This section provides useful insights into when to use pointers or values based on performance considerations in Go,
//especially for large data structures. Let’s summarize the key points and provide an example to test these behaviors.
//Key Takeaways
//
//Input Parameters:
//Pointer: Performance is constant (~1 nanosecond), regardless of data size.
//Value: Performance degrades with increasing data size, especially for large structs.
//
//Return Values:
//Small structs (< 1MB): Returning a value is faster than a pointer.
//Large structs (> 1MB): Returning a pointer is faster than a value.
//
//Practical Considerations:
//For most real-world use cases, the difference is negligible unless you're working with extremely large data structures.
//When handling immutable large data, pointers can be an efficient choice.
