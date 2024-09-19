package main

import "math"
import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Define the Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}

	// Call PrintArea with different shapes
	PrintArea(circle)     // Output: Area: 78.54
	PrintArea(rectangle)  // Output: Area: 50.00
}
