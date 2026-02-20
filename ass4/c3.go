
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

type ColoredShape interface {
	Shape
	Color() string
}

type ColoredCircle struct {
	Circle
	color string
}

func (cc ColoredCircle) Color() string {
	return cc.color
}

func main() {
	
	cc := ColoredCircle{
		Circle: Circle{radius: 5},
		color:  "Red",
	}

	
	fmt.Printf("Area of ColoredCircle: %.2f\n", cc.Area())
	
	fmt.Printf("Color of ColoredCircle: %s\n", cc.Color())
}		