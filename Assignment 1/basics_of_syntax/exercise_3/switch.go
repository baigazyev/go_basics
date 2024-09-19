package main

import (
	"fmt"
)

func main() {
	var day int
	fmt.Print("Enter a number (1 for Monday, 2 for Tuesday, etc.): ")
	fmt.Scan(&day)

	// Using a switch statement to print the day of the week
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid input! Please enter a number between 1 and 7.")
	}
}
