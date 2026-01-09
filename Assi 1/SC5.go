package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first string: ")
	str1, _ := reader.ReadString('\n')
	str1 = strings.TrimSpace(str1)

	fmt.Print("Enter second string: ")
	str2, _ := reader.ReadString('\n')
	str2 = strings.TrimSpace(str2)

	if strings.Contains(str1, str2) {
		fmt.Println(str2, "is a substring of", str1)
	} else {
		fmt.Println(str2, "is not a substring of", str1)
	}
}
