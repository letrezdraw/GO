// write a go program that creates a slice of integers, check numbers from slice are even or odd and further sent to resrctive go routines throuht channels and display values recived by go routines.

package main

import (
	"fmt"
)

func checkEvenOdd(num int, evenCh chan int, oddCh chan int) {
	if num%2 == 0 {
		evenCh <- num
	} else {
		oddCh <- num
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenCh := make(chan int)
	oddCh := make(chan int)

	for _, num := range numbers {
		go checkEvenOdd(num, evenCh, oddCh)
	}

	for i := 0; i < len(numbers); i++ {
		select {
		case evenNum := <-evenCh:
			fmt.Printf("Even number: %d\n", evenNum)
		case oddNum := <-oddCh:
			fmt.Printf("Odd number: %d\n", oddNum)
		}
	}
}
