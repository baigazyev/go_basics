package main

import (
	"fmt"
)

// Function to swap two strings
func swap(s1, s2 string) (string, string) {
	return s2, s1
}

func main() {
	first, second := swap("Hello", "World")
	fmt.Println("Swapped strings:", first, second)  // Output: Swapped strings: World Hello
}
