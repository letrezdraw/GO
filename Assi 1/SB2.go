package main

import "fmt"

func main() {
	rows := 5

	for i := 0; i < rows; i++ {
		for j := 0; j < rows-i-1; j++ {
			fmt.Print(" ")
		}
		num := 1
		for j := 0; j <= i; j++ {
			fmt.Print(num, " ")
			num = num * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}
