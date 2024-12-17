package main

//A pointer is simply a variable that holds the location in memory where a
//value is stored. If you’ve taken computer science courses, you might have
//seen a graphic to represent how variables are stored in memory.
//The zero value for a pointer is nil. We’ve seen nil a few times before, as
//the zero value for slices, maps, and functions. All of these types are
//implemented with pointers. (Two more types, channels and interfaces, are also implemented with pointers).

import "fmt"

// Helper function to return a pointer to a string
func stringp(s string) *string {
	return &s
}

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func main() {
	// Basic pointer usage
	x := 10
	pointerToX := &x
	fmt.Println("Pointer to x:", pointerToX)                // prints memory address
	fmt.Println("Value of x through pointer:", *pointerToX) // prints 10

	z := 5 + *pointerToX
	fmt.Println("z:", z) // prints 15

	// Nil pointer
	var nilPointer *int
	fmt.Println("Is nilPointer nil?", nilPointer == nil) // prints true
	// Uncommenting the next line would cause a runtime panic
	// fmt.Println(*nilPointer)  // panic: runtime error: invalid memory address or nil pointer dereference

	// Using new function
	var y = new(int)
	fmt.Println("Is y nil?", y == nil) // prints false
	fmt.Println("Value of y:", *y)     // prints 0 (new initializes it to zero)

	// Using pointer with structs
	var middleName = "Perry"
	p := person{
		FirstName:  "Pat",
		MiddleName: stringp(middleName), // Using the helper function
		LastName:   "Peterson",
	}
	fmt.Println("Person:", p)

	// Trying to directly assign a literal to a pointer field (will cause a compile-time error)
	// p2 := person{
	//     FirstName: "John",
	//     MiddleName: "Doe", // This will not compile
	//     LastName: "Smith",
	// }

	// Correct way to assign a pointer to a string literal
	var firstName = "John"
	p2 := person{
		FirstName:  firstName,
		MiddleName: stringp("Doe"),
		LastName:   "Smith",
	}
	fmt.Println("Another Person:", p2)
}
