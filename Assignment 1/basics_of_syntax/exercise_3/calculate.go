package main

import (
	"fmt"
)

func main() {
	sum := 0
	// Using a for loop to calculate the sum
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("The sum of the first 10 natural numbers is:", sum)
}
