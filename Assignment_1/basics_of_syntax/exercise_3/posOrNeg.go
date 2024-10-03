package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&num)

	// Using if-else statements
	if num > 0 {
		fmt.Println("The number is positive.")
	} else if num < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}
}
