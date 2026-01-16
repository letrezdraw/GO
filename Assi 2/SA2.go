package main

import "fmt"

func main() {
	var num, sum int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	sum = sumOfDigits(num)
	fmt.Printf("Sum of digits of %d is %d\n", num, sum)
}
func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + sumOfDigits(n/10)
}
