package main

import "fmt"

// Define a struct Foo
type Foo struct {
	Name string
	Age  int
}

func main() {
	// Creating a pointer to Foo with zero values
	x := &Foo{}

	// Assigning values to the struct fields
	x.Name = "John"
	x.Age = 30

	fmt.Println(x)  // Prints: &{John 30}
	fmt.Println(*x) // Prints: {John 30}
}
