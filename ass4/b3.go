//copy all elements of one array to another array.using method
package main

import "fmt"

func copyArray(src []int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	return dest
}

func main() {
	src := []int{1, 2, 3, 4, 5}
	dest := copyArray(src)

	fmt.Println("Source array:", src)
	fmt.Println("Destination array:", dest)
}	