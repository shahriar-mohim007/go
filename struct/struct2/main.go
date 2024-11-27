package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	// Use pointers to structs
	people := []*Person{
		{Name: "Alice", Age: 30, City: "New York"},
		{Name: "Bob", Age: 25, City: "Los Angeles"},
	}

	// Modify the values directly using the pointer
	for _, person := range people {
		person.Name = "mohim" // Modify the original struct via pointer
	}

	// Print the modified slice
	for _, person := range people {
		fmt.Println("Name:", person.Name)
		fmt.Println("Age:", person.Age)
		fmt.Println("City:", person.City)
		fmt.Println()
	}
}
