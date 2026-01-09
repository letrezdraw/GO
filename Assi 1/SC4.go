package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num >= 0 && num <= 9 {
		fmt.Println(num, "is a single digit")
	} else {
		fmt.Println(num, "is not a single digit")
	}
}
