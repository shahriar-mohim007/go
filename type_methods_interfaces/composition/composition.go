package main

import (
	"fmt"
)

// Employee struct represents an individual employee with a Name and ID.
type Employee struct {
	Name string
	ID   string
}

// Description method provides a string representation of the Employee.
func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

// Manager struct embeds Employee and includes a list of direct reports.
type Manager struct {
	Employee
	Reports []Employee
}

// FindNewEmployees is a stubbed method for business logic to find new employees.
func (m Manager) FindNewEmployees() []Employee {
	// Placeholder for actual business logic
	return []Employee{
		{Name: "Alice Doe", ID: "67890"},
		{Name: "Charlie Smith", ID: "11223"},
	}
}

type Inner struct {
	X int
}
type Outer struct {
	Inner
	X int
}

func main() {
	// Creating a Manager instance
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}

	// Accessing promoted fields and methods from the embedded Employee
	fmt.Println(m.ID)            // Prints: 12345
	fmt.Println(m.Description()) // Prints: Bob Bobson (12345)

	// Using the FindNewEmployees method
	newEmployees := m.FindNewEmployees()
	fmt.Println("New Employees:")
	for _, emp := range newEmployees {
		fmt.Println(emp.Description())
	}

	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)
	// prints 20
	fmt.Println(o.Inner.X)
}
