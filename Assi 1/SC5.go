package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Print("Enter first string: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter second string: ")
	fmt.Scanln(&str2)

	if strings.Contains(str2, str1) {
		fmt.Println(str1, "is a substring of", str2)
	} else {
		fmt.Println(str1, "is not a substring of", str2)
	}
}
