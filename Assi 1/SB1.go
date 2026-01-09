package main

import "fmt"

func main() {
	var t int = 4
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", t, i, t*i)
	}
}
