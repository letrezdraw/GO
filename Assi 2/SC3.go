package main

import (
	"errors"
	"fmt"
)

func f(a int) (int, error) {
	if a == 0 {
		return 0, errors.New("Zero not allowed")
	} else {
		return a * 2, nil
	}
}

func main() {
	v, e := f(0)
	if e != nil {
		fmt.Println("Error:", e)
	} else {
		fmt.Println("Value:", v)
	}
}
