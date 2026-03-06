package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(id int) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
	}
}

func main() {
	for i := 1; i <= 5; i++ {
		go generateNumbers(i)
	}

	time.Sleep(3 * time.Second)
}