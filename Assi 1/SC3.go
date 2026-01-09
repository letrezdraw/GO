package main

import "fmt"

func main() {
	a := 15
	b := 5
	operation := "+"

	switch operation {
	case "+":
		fmt.Println(a, "+", b, "=", a+b)
	case "-":
		fmt.Println(a, "-", b, "=", a-b)
	case "*":
		fmt.Println(a, "*", b, "=", a*b)
	case "/":
		fmt.Println(a, "/", b, "=", a/b)
	default:
		fmt.Println("Invalid operation")
	}
}
