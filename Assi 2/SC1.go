package main

import "fmt"

func callbyvalue(a int) {
	a = a + 10
	fmt.Println("Inside callbyvalue:", a)
}

func main() {
	x := 10
	callbyvalue(x)
	fmt.Println("Inside main:", x)
}
