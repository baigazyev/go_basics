package main

import (
	"fmt"
)



type Person struct {
	Name string
	Age  int
}


func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	// Create an instance of Person and set its fields
	person := Person{
		Name: "Alice",
		Age:  25,
	}

	// Call the Greet method on the person instance
	person.Greet()  // Output: Hello, my name is Alice and I am 25 years old.
}
