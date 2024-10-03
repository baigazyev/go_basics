package main

import (
	"fmt"
)

// Define the Employee struct
type Employee struct {
	Name string
	ID   int
}

// Define the Manager struct which embeds Employee and adds a Department field
type Manager struct {
	Employee   // Embedding Employee
	Department string
}

// Define the Work method for Employee
func (e Employee) Work() {
	fmt.Printf("Employee Name: %s, ID: %d\n", e.Name, e.ID)
}

func main() {
	// Create an instance of Manager and set its fields
	manager := Manager{
		Employee: Employee{
			Name: "Alice",
			ID:   101,
		},
		Department: "Engineering",
	}

	// Call the Work method on the Manager instance
	manager.Work()  // Output: Employee Name: Alice, ID: 101
}
