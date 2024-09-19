package main

import (
	"fmt"
)

// Function to add two integers
func add(a int, b int) int {
	return a + b
}

func main() {
	sum := add(5, 10)
	fmt.Println("Sum:", sum)  // Output: Sum: 15
}
