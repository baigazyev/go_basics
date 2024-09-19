package main

import (
	"fmt"
)

func main() {
	// Using var keyword
	var age int = 25
	var height float64 = 5.9
	var name string = "John Doe"
	var isStudent bool = true

	// Using short declaration syntax
	country := "Kazakhstan"
	gpa := 3.8
	isEmployed := false

	// Printing values and types of variables using fmt.Printf
	fmt.Printf("Age: %d, Type: %T\n", age, age)
	fmt.Printf("Height: %.1f, Type: %T\n", height, height)
	fmt.Printf("Name: %s, Type: %T\n", name, name)
	fmt.Printf("Is student: %t, Type: %T\n", isStudent, isStudent)
	fmt.Printf("Country: %s, Type: %T\n", country, country)
	fmt.Printf("GPA: %.1f, Type: %T\n", gpa, gpa)
	fmt.Printf("Is employed: %t, Type: %T\n", isEmployed, isEmployed)
}
