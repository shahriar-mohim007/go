package main

import "fmt"

func main() {
	// Anonymous struct with fields: name, age, and pet
	var person struct {
		name string
		age  int
		pet  string
	}
	// Assign values to the anonymous struct
	person.name = "Bob"
	person.age = 50
	person.pet = "Dog"

	// Print values of the anonymous struct
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
	fmt.Println("Pet:", person.pet)

	// Another example of an anonymous struct used for a pet
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "Dog",
	}

	// Print values of the second anonymous struct (pet)
	fmt.Println("\nPet name:", pet.name)
	fmt.Println("Pet kind:", pet.kind)
}
