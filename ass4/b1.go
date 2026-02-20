package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (s *Student) show() {
	fmt.Printf("Student Name: %s, Age: %d\n", s.name, s.age)
}
func main() {
	s := Student{name: "Soham Gajare", age: 20}
	s.show()
}
