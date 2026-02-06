//create and print multidimensional slice
package main

import "fmt"

func main() {
	multiSlice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Multidimensional Slice elements are:", multiSlice)
}