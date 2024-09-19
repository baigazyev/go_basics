package main

import (
	"fmt"
)

// Function to return quotient and remainder
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

func main() {
	quot, rem := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", quot, rem)  // Output: Quotient: 3, Remainder: 1
}
