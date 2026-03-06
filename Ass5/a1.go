//wap a go program routine and channel that will print the sum of the squares and cubes of the individual digits of a number.

package main

import "fmt"

func sumOfSquaresAndCubes(n int, ch chan int) {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit*digit + digit*digit*digit
		n /= 10
	}
	ch <- sum
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	ch := make(chan int)
	go sumOfSquaresAndCubes(num, ch)

	result := <-ch
	fmt.Printf("The sum of squares and cubes of the digits of %d is: %d\n", num, result)
}
