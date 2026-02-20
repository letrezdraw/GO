
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	var s Shape

	s = Circle{radius: 5}
	fmt.Printf("Area of Circle: %.2f\n", s.Area())

	s = Rectangle{width: 4, height: 6}
	fmt.Printf("Area of Rectangle: %.2f\n", s.Area())
}