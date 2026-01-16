package main

import "fmt"

func isPalindrome(n int) bool {
	original := n
	reversed := 0
	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}
	return original == reversed
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	if isPalindrome(num) {
		fmt.Printf("%d is a palindrome number", num)
	} else {
		fmt.Printf("%d is not a palindrome number", num)
	}
}
