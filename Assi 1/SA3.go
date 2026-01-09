package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Before swap: a = %d, b = %d", a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("After swap: a = %d, b = %d", a, b)
}