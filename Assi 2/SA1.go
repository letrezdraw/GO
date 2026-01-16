package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	n := add(5, 10)
	fmt.Println(n)
}
