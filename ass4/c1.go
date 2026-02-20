//create interface and display its value with help of type assertion

package main

import "fmt"

// Shape interface with Area method
type Shape interface {
	Area() float64
}

// Circle struct with radius
type Circle struct {
	radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var s Shape

	s = Circle{radius: 5}
	fmt.Printf("Area of Circle: %.2f\n", s.Area())

	// Type assertion to access Circle-specific fields
	if circle, ok := s.(Circle); ok {
		fmt.Printf("Circle Radius: %.2f\n", circle.radius)
	} else {
		fmt.Println("s is not a Circle")
	}
}	
