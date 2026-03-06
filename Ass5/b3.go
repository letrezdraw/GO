//wap how to create channel and illustrate how to close a channel using for range loop and close function.
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch { 
		fmt.Println(val)
	}
}
