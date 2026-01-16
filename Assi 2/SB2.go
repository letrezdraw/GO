package main

import "fmt"

func try(b int) (x,y int) {
	x = b + 1
	y = b + 5
	return
}

func main() {
	a, b := try(3)
	fmt.Println(a, b)
}