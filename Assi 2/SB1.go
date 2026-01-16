package main

import "fmt"

  func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Before swapping: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swapping: a = %d, b = %d\n", a, b)
}
