// wap main go routine to read and write fibonacci seires to the channel.
package main

import "fmt"

func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	n := 10
	ch := make(chan int)

	go fibonacci(n, ch)

	fmt.Printf("Fibonacci series up to %d terms:\n", n)
	for num := range ch {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
