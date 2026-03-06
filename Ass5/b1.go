//wap in go to create buffered channel, store few values in it and find channel capacity and length.read values from channel and find modified length of a channel.

package main

import "fmt"

func main() {
	ch := make(chan int, 5)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Printf("Channel capacity: %d\n", cap(ch))
	fmt.Printf("Channel length: %d\n", len(ch))

	fmt.Printf("Read value: %d\n", <-ch)
	fmt.Printf("Read value: %d\n", <-ch)

	fmt.Printf("Modified channel length: %d\n", len(ch))
}
