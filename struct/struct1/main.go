package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	people := []Person{
		{Name: "Alice", Age: 30, City: "New York"},
		{Name: "Bob", Age: 25, City: "Los Angeles"},
	}

	// Modify the actual value by using the index
	for i := range people {
		people[i].Name = "mohim" // Modify the original value in the slice
	}

	// Print the modified slice
	for _, person := range people {
		fmt.Println("Name:", person.Name)
		fmt.Println("Age:", person.Age)
		fmt.Println("City:", person.City)
		fmt.Println()
	}
}
