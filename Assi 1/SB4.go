package main

import "fmt"

func main() {
	x := 10
	ptr := &x
	ptrToPtr := &ptr

	fmt.Println("x =", x)
	fmt.Println("*ptr =", *ptr)
	fmt.Println("**ptrToPtr =", **ptrToPtr)

	**ptrToPtr = 20
	fmt.Println("After change: x =", x)
}
