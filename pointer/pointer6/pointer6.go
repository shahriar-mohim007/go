package main

import "fmt"

func modifySlice(s []int) {
	s[0] = 42 // Modify content
}

func appendToSlice(s []int) {
	s = append(s, 40)                       // Append changes the slice but doesn't reflect in the original
	fmt.Println("Inside appendToSlice:", s) // Output: [10 20 30 40]
}

func appendToSlice1(s *[]int) {
	*s = append(*s, 40) // Modify the original slice through the pointer
}

func main() {
	mySlice := []int{10, 20, 30}
	fmt.Println(cap(mySlice), len(mySlice))
	fmt.Println("Before modifySlice:", mySlice) // Output: [10 20 30]
	modifySlice(mySlice)
	fmt.Println("After modifySlice:", mySlice) // Output: [42 20 30]

	appendToSlice(mySlice)
	fmt.Println("After appendToSlice:", mySlice) // Output: [10 20 30]
	appendToSlice1(&mySlice)                     // Pass a pointer to the slice
	fmt.Println("After appendToSlice:", mySlice) // Output: [10 20 30 40]

}

//The appendToSlice function does not reflect changes in the original mySlice because of how
//slices are passed to functions in Go and how append works. Let me explain:
//Slices Are Passed by Value
//
//When you pass a slice to a function, Go passes a copy of the slice descriptor
//(which contains the length, capacity, and pointer to the underlying array).
//	This means the function operates on a copy of the slice descriptor, not the original.
//
//Modifications to the underlying array (e.g., changing s[0]) are reflected in the original
//because the copy of the descriptor points to the same memory.
//Changes to the slice itself (e.g., appending) do not affect the original slice outside the
//function because the descriptor is modified in isolation.

/*This excerpt provides a comprehensive explanation of memory allocation and garbage collection in Go,
focusing on the practical and performance aspects. Here's a summary of the key points covered:
Garbage Collection in Go

Definition of Garbage: Data without pointers pointing to it, which can be reclaimed for reuse.
Purpose: Prevent memory leaks by automatically identifying and recovering unused memory.
Performance Considerations:
Go's garbage collector prioritizes low latency, ensuring each collection cycle is completed in less than 500 microseconds.

Heap vs. Stack
Stack Allocation

Characteristics:
Fast and simple due to its sequential allocation using a stack pointer.
Stores local variables, function parameters, and return values.
Memory is automatically reclaimed when the function exits.
Advantages in Go:
Each goroutine has its own stack, allowing dynamic resizing during execution.
Limitations:
Requires size to be known at compile time.
Local variables must not escape their function to stay on the stack.

Heap Allocation

Characteristics:
Used for data that can't be stored on the stack (e.g., unknown sizes, data escaping function scope).
Managed by the garbage collector.
Performance Impact:
Slower due to non-sequential memory access and garbage collection overhead.
Data scattered across RAM, reducing read efficiency.

Escape Analysis in Go

What It Does:
Determines if variables can be safely allocated on the stack.
Decisions:
If conditions for stack storage aren't met, data escapes to the heap.
Compiler Conservatism:
Prioritizes safety to avoid memory corruption.
Improvements:
Go's escape analysis has become more precise with newer versions.

Optimizing Garbage Collection

Minimize Heap Allocations:
Use value types (e.g., structs) over pointers when possible.
Allocate memory for slices of structs or primitive types sequentially for better cache performance.
Avoid Creating Excessive Garbage:
Go idioms and best practices inherently encourage efficient memory use.*/
//Summary of Memory Storage:
//
//    Stack: Local variables in the main function (e.g., x := 42) and pointers to dynamically allocated objects.
//    Heap: Dynamically allocated objects (e.g., new(Point)), if they are created inside main and need to persist beyond the function.
//    Data Section: Global variables declared outside any function but within the main package.
