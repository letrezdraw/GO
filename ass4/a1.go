
//go code to create an interface shape that includes area and perimeter.implement in circle and rectangle

package main

import (
	"fmt"
	"math"
)

// Shape interface with Area and Perimeter methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct with radius
type Circle struct {
	radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Rectangle struct with width and height
type Rectangle struct {
	width  float64
	height float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 6}

	fmt.Printf("Circle Area: %.2f\n", circle.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", circle.Perimeter())

	fmt.Printf("Rectangle Area: %.2f\n", rectangle.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rectangle.Perimeter())
}		